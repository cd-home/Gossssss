package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {

}

func ClientUploadFile() {
	url := "http://127.0.0.1:8080/upload"
	body := new(bytes.Buffer)
	bw := multipart.NewWriter(body)
	part, err := bw.CreateFormFile("upload", "backup.html")
	if err != nil {
		log.Fatal(err)
	}
	src, err := os.Open("../static/index.html")
	if err != nil {
		log.Fatal(err)
	}
	// data
	_, err = io.Copy(part, src)
	if err != nil {
		log.Fatal(err)
	}
	// fields
	err = bw.WriteField("name", "yao")
	if err != nil {
		log.Fatal(err)
	}
	err = bw.Close()
	if err != nil {
		log.Fatal(err)
	}
	response, err := http.Post(url, bw.FormDataContentType(), body)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(response.Status)
	}
}
