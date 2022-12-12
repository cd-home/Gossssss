package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func main() {

}

func SimpleGet() {
	url := "https://www.baidu.com"
	// ?k=v if u want to
	response, err := http.Get(url)
	if err != nil {
		return
	}
	log.Println(response.Status)
}

func SimplePost() {
	url := ""
	contentType := "application/json"
	data := struct {
		Name string
		Age  int
	}{
		Name: "yao",
		Age:  27,
	}
	bs, _ := json.Marshal(data)
	http.Post(url, contentType, bytes.NewReader(bs))
}
