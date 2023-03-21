package main

import (
	"crypto/tls"
	"fmt"
	"strings"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
	"log"
	"encoding/json"
	"os"
)
type sessionCookie struct {
	Ipaddress string
	Phpsessid string
	Time      time.Time
}
/*
type h struct {
	H []sessionCookie
}*/
// The functions APISessionAuth(...) and getAPIData(...) utilizes curl-to-go translator but is modified for cookie management.
// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

// curl -k --data "Username=student&Password=PanneKake23" -i -v https://10.233.230.11/rest/login

// TODO: This is insecure; use only in dev environments.
func APISessionAuth(username string, password string, ipaddress string) (string,error) {
	//var read []byte
	var phpsessid string
	read,err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println("No file yet")
	}
	//struct := &Host{}
	//var str Name
	//doc := make(map[string]Host{})
	Hosts := []sessionCookie{}
	err = json.Unmarshal(read, &Hosts)
	if err != nil {
		fmt.Println("No data retrieved unmarhalling json phpsessid")
	} else {
//Checks if current time is 8 min after logged time in json file
//If so, use the sessioncookie stored in the json file
		for i :=range Hosts {
			if (time.Now().Before(Hosts[i].Time.Add(2 * time.Minute))){ //Hosts.Time.After(time.Now().Add(1 * time.Minute))) {
				if (Hosts[i].Ipaddress == ipaddress) {
					fmt.Println("retrieved from file")
					phpsessid = Hosts[i].Phpsessid
			}
		}
	}
	return phpsessid,nil
}


	cfg := getConf(&Config{})
	timeout := cfg.Authtimeout
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr,Timeout: time.Duration(timeout) * time.Second}

	params := url.Values{}
	params.Add("Username", username)
	params.Add("Password", password)
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", "https://"+ipaddress+"/rest/login", body)
	if err != nil {
		log.Flags()
			fmt.Println("error in auth:", err)
			return "Error fetching data", err
		//	fmt.Println("error in systemExporter:", error)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		log.Flags()
		fmt.Println("error in auth:", err)
		return "Error fetching data", err
		//fmt.Println("error in systemExporter:", err)
	}

	  m := make(map[string]string)
	  for _, c := range resp.Cookies() {
		 m[c.Name] = c.Value
	  }
	 // fmt.Println(m["PHPSESSID"])
	phpsessid = m["PHPSESSID"]

	data := sessionCookie{ipaddress, phpsessid, time.Now()}
	jsonByte, _ := json.Marshal(data)

	f, err := os.OpenFile("./data.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return "Cannot open file",err
	}
	defer f.Close()

	n, err := f.Write(jsonByte)
	if err != nil {
		fmt.Println(n, err)
	}

	if n, err = f.WriteString("\n"); err != nil {
		fmt.Println(n, err)
	}

	defer resp.Body.Close()


return phpsessid,nil
}


func getAPIData(url string, phpsessid string) (string,error){

tr2 := &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
}
client2 := &http.Client{Transport: tr2}
cookie1 := &http.Cookie{
	Name:   "PHPSESSID",
	Value:  phpsessid,
	//Path:     "/",
	MaxAge:   3600,
	HttpOnly: false,
	Secure:   true,
}
req2, err := http.NewRequest("GET", url, nil)
if err != nil {
	log.Flags()
		fmt.Println("error in getapidata():", err)
		return "Error fetching data", err
	//	fmt.Println("error in systemExporter:", error)
}
req2.AddCookie(cookie1)
	resp2, err := client2.Do(req2)
	if err != nil {
		log.Flags()
			fmt.Println("error in getapidata():", err)
			return "Error fetching data", err
	}

	b, err := ioutil.ReadAll(resp2.Body)
	defer resp2.Body.Close()

	return string(b), err
}
/*
func () {
	php, err  := APISessionAuth("student", "PanneKake23", "10.233.234.11")

	fmt.Println(php,err)
}
*/