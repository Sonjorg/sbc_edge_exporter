package main

import (
	"crypto/tls"
	"fmt"
	"strings"
	//"github.com/tiket-oss/phpsessgo"
	//"io/ioutil"
	"net/http"
	//"net/http/cookiejar"
	//"net/http/cookiejar"
	"net/url"
	//"strconv"

)

// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

// curl -k --data "Username=student&Password=PanneKake23" -i -v https://10.233.230.11/rest/login

// TODO: This is insecure; use only in dev environments.
func main() {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	params := url.Values{}
	params.Add("Username", `student`)
	params.Add("Password", `PanneKake23`)
	body := strings.NewReader(params.Encode())

	req, err := http.NewRequest("POST", "https://10.233.230.11/rest/login", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		// handle err
	}
	
	for _, cookie := range resp.Cookies() {
		fmt.Println("Found a cookie named:", cookie.Name, cookie.Value, cookie.String())
	  }

	/*func (e *PHPSessionEncoder) Decode(raw string) (phpencode.PhpSession, error) {
		resp.Cookies().string()
	}*/
	//str := strconv.Itoa(resp.Cookies())
	/*
	fmt.Println("HER KOMMER COOKIES ", resp.Cookies())*/
	defer resp.Body.Close()
	///////////////////////////////////////////////////////////////////////
	// new request with cookie from authentication
	// https://golangbyexample.com/set-cookie-http-golang/

	//string := resp.Cookies().token
	/*cookie := &http.Cookie{
		Name:   "token",
		Value:  resp.Cookies().string(),
		MaxAge: 300,
	}*/

	// curl --cookie SkrivInnCookieHer \ -i -k https://10.233.230.11/rest/isdnsg/10001

	// TODO: This is insecure; use only in dev environments.
	/*
	tr2 := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client2 := &http.Client{Transport: tr2}

	req2, err := http.NewRequest("GET", "https://10.233.230.11/rest/isdnsg/10001", nil)
	if err != nil {
		// handle err
	}
	
	//req2.ph
	//legger ved cookie før request
	//req2.AddCookie()
	//url := "https://10.233.230.11/rest/isdnsg/10001"
	//u1, _ :=
	client2.Jar.SetCookies(req2.URL, resp.Cookies())

	resp2, err := client2.Do(req2)
	if err != nil {
		// handle err
	}


	println(resp2)
	defer resp2.Body.Close()
*/
}
