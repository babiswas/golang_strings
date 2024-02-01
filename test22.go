package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://learningmanager.adobe.com/primeapi/v2/badges", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Setting Headers:")
	req.Header = http.Header{
		"Accept":        {"application/vnd.api+json"},
		"Content-Type":  {"application/json"},
		"Authorization": {"oauth " + "c884da4b41a67c69062f4a105a4c8935"},
	}
	q := req.URL.Query()
	q.Add("sort", "name")
	q.Add("page[offset]", "0")
	q.Add("page[limit]", "1")
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Status)
	fmt.Println(string(responseBody))
}
