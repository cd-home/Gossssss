package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

//func TestSimpleGetByDefaultClient(t *testing.T) {
//	url := "https://www.baidu.com"
//	response, err := http.Get(url)
//	if err != nil {
//		return
//	}
//	t.Log(response.Status)
//}

func main() {
	url := "http://127.0.0.1:8080/upload"
	body := new(bytes.Buffer)
	bw := multipart.NewWriter(body)
	part, err := bw.CreateFormFile("upload", "backup.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	src, err := os.Open("../static/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	// data
	_, err = io.Copy(part, src)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = bw.Close()
	// fields
	err = bw.WriteField("name", "yao")
	if err != nil {
		fmt.Println(err)
		return
	}

	response, err := http.Post(url, bw.FormDataContentType(), body)
	if err != nil {
		return
	} else {
		fmt.Println(response)
	}
}
