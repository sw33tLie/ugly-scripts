package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var username, password, grafanaHost string
	flag.StringVar(&username, "U", "", "admin")
	flag.StringVar(&password, "P", "", "admin")
	flag.StringVar(&grafanaHost, "h", "", "https://grafana.example.com:3000")
	flag.Parse()

	login(username, password, grafanaHost)
}

func login(username string, password string, url string) {

	var jsonStr = []byte(`{"user":"` + username + `","password":"` + password + `","email":""}`)
	req, err := http.NewRequest("POST", url+"/login", bytes.NewBuffer(jsonStr))
	req.Header.Set("Origin", url+"/login")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:75.0) Gecko/20100101 Firefox/75.0")
	req.Header.Set("Referer", url+"/signup")
	req.Header.Set("Connection", "close")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response Body:", string(body))
}
