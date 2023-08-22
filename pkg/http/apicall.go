/* Copyright (C) 2023 Sondre Jørgensen - All Rights Reserved
 * You may use, distribute and modify this code under the
 * terms of the CC BY 4.0 license
*/
package http

import (
	"edge_exporter/pkg/database"
	"edge_exporter/pkg/config"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"io"
	"strings"
	"time"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)


func SBCIsDown(ipaddress string) bool{
	cfg := config.GetConfig(&config.HostCompose{})
	timeout := cfg.Authtimeout
	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
	req, err := http.NewRequest("GET", "https://"+ipaddress+"/", nil)
	if err != nil {
			log.Println("1", err)
	}
	_, err = client.Do(req)
	if err != nil {
			log.Println("2", err)
	}
	res := strings.Contains(err.Error(), "certificate")
	fmt.Println(res)
	return !res
}
// The functions APISessionAuth(...) and getAPIData(...) utilizes curl-to-go translator but is modified for cookie management and storing these with sqlite.
// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
// curl -k --data "Username=un&Password=pw" -i -v https://ipadr/rest/login
// TODO: This is insecure; use only in dev environments.
func APISessionAuth(username string, password string, ipaddress string) (string, error) {
	var phpsessid string
	var err error

	var sqliteDatabase *sql.DB

	sqliteDatabase, err = sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		log.Print(err)
	}
	defer sqliteDatabase.Close()
//Trying to fetch session cookie from database if not expired
	phpsessid, err = database.GetSqliteKeyIfNotExpired(sqliteDatabase, ipaddress)
	if phpsessid != "" {
		return phpsessid, nil
	}

	cfg := config.GetConf(&config.Config{})
	timeout := cfg.Authtimeout
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr, Timeout: time.Duration(timeout) * time.Second}

	params := url.Values{}
	params.Add("Username", username)
	params.Add("Password", password)
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", "https://"+ipaddress+"/rest/login", body)
	if err != nil {
		log.Flags()
		log.Print("error authentication (APISessionAuth):", err)
		return "Error fetching data", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		//log.Flags()
		log.Print("error in auth:", err)
		return "Error fetching data", err
	}

	m := make(map[string]string)
	for _, c := range resp.Cookies() {
		m[c.Name] = c.Value
	}
	phpsessid = m["PHPSESSID"]

	defer resp.Body.Close()

	now := time.Now().Format(time.RFC3339)
	//As session cookie is fetched from http, inserting it into database
	if (database.RowExists(sqliteDatabase,ipaddress)) {
		database.Update(sqliteDatabase,phpsessid,now,ipaddress)
	} else {
		database.InsertAuth(sqliteDatabase, ipaddress, phpsessid, now)
	}
	return phpsessid, nil
}



//Used to fetch all data except session cookies and chassis labels
func GetAPIData(url string, phpsessid string) (string, []byte, error) {

	tr2 := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client2 := &http.Client{Transport: tr2}
	cookie1 := &http.Cookie{
		Name:  "PHPSESSID",
		Value: phpsessid,
		//Path:     "/",
		MaxAge:   3600,
		HttpOnly: false,
		Secure:   true,
	}
	req2, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Flags()
		log.Print("error in getapidata():", err)
		return "Error fetching data", nil, err
	}
	req2.AddCookie(cookie1)
	resp2, err := client2.Do(req2)
	if err != nil {
		log.Flags()
		log.Print("error in getapidata():", err)
		return "Error fetching data", nil, err
	}

	b, err := io.ReadAll(resp2.Body)
	defer resp2.Body.Close()

	return "Success", b, err
}
