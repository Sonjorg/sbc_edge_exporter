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
	//"regexp"
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
	
	  m := make(map[string]string)
	  for _, c := range resp.Cookies() {
		 m[c.Name] = c.Value
	  }
	  fmt.Println(m["PHPSESSID"])
	  phpsessid := m["PHPSESSID"]
	  
	  //var exp = regexp.MustCompile(`PHPSESSID=.+;`)
	  //matchType := exp.FindStringSubmatch(c2)
	 // fmt.Println(matchType)
	  
//value := cookie.Value
//fmt.Println(c)
defer resp.Body.Close()
	//str := strconv.Itoa(resp.Cookies())
	/*
	fmt.Println("HER KOMMER COOKIES ", resp.Cookies())*/
	///////////////////////////////////////////////////////////////////////
	// new request with cookie from authentication
	// https://golangbyexample.com/set-cookie-http-golang/

	//string := resp.Cookies().token
	/*cookie := &http.Cookie{
		Name:   "PHPSESSID",
		Value:  c,
		MaxAge: 300,
	}*/


	// TODO: This is insecure; use only in dev environments.

	/*tr2 := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	} */

	/*
	jar, err := cookiejar.New(nil)
    if err != nil {
      fmt.Println(err)//  log.Fatalf("Got error while creating cookie jar %s", err.Error())
    }
	var client2 http.Client

    client2 = http.Client{
        Jar: jar,
    }

	cookie1 := &http.Cookie{
        Name:   "PHPSESSID",
        Value:  phpsessid,
		Path:     "/",
        MaxAge:   3600,
        HttpOnly: true,
        Secure:   true,
    }


	req2, err := http.NewRequest("GET", "https://10.233.230.11/rest/isdnsg/10001", nil)
	if err != nil {
		fmt.Println(err)
	}

	//url := "https://10.233.230.11/rest/isdnsg/10001"
	//u1, _ :=
	//client2.Jar.SetCookies(req2.URL, c)
	req2.AddCookie(cookie1)//SetCookie(cookie1)
	resp2, err := client2.Do(req2)
	if err != nil {
		fmt.Println(err)
	}
	println(resp2.Body)
	resp2.Body.Close()
*/

// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

// curl --cookie  \ -i -k https://10.233.230.11/rest/isdnsg/10001

// TODO: This is insecure; use only in dev environments.
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
req2, err := http.NewRequest("GET", "https://10.233.230.11/rest/isdnsg/10001", nil)
if err != nil {
	// handle err
}
req2.AddCookie(cookie1)
resp2, err := client2.Do(req2)
if err != nil {
	// handle err
}
fmt.Println(resp2)
defer resp2.Body.Close()

}
