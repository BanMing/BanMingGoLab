package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

var wc sync.WaitGroup

func HelloClient(client *http.Client, url string, method string) {
	defer wc.Done()
	var reqData io.Reader = nil
	if method == "POST" {
		reqData = strings.NewReader("name=ali&age=19")
	}

	req, err := http.NewRequest(method, url, reqData)
	if err != nil {
		log.Fatal(err)
	}
	//设置Content-Type很重要
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rep, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(rep.Body)
	rep.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", data)
}

func main() {
	client := &http.Client{}
	rootURL := "http://127.0.0.1:9999/"
	wc.Add(1)
	go HelloClient(client, rootURL+string("login"), "GET")
	wc.Add(1)
	go HelloClient(client, rootURL+string("post"), "POST")
	wc.Wait()
}
