package client_test

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"testing"
)

func TestSimpleGetByDefaultClient(t *testing.T) {
	url := "https://www.baidu.com"
	// ?k=v if u want to
	response, err := http.Get(url)
	if err != nil {
		return
	}
	t.Log(response.Status)
}

func TestHttpPostFile(t *testing.T) {
	url := "http://127.0.0.1:8080/upload"
	body := new(bytes.Buffer)
	bw := multipart.NewWriter(body)
	part, err := bw.CreateFormFile("upload", "backup.html")
	if err != nil {
		t.Error(err)
	}

	src, err := os.Open("../static/index.html")
	if err != nil {
		t.Error(err)
	}
	// data
	_, err = io.Copy(part, src)
	if err != nil {
		t.Error(err)
	}

	// fields
	err = bw.WriteField("name", "yao")
	if err != nil {
		t.Error(err)
	}

	err = bw.Close()

	if err != nil {
		fmt.Println(err)
	}

	response, err := http.Post(url, bw.FormDataContentType(), body)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(response.Status)
	}
}

func TestPipe(t *testing.T) {
	name := ""
	url := ""
	r, w := io.Pipe()
	m := multipart.NewWriter(w)
	go func() {
		defer w.Close()
		defer m.Close()
		part, err := m.CreateFormFile("myFile", "foo.txt")
		if err != nil {
			return
		}
		file, err := os.Open(name)
		if err != nil {
			return
		}
		defer file.Close()
		if _, err = io.Copy(part, file); err != nil {
			return
		}
	}()
	http.Post(url, m.FormDataContentType(), r)
}
