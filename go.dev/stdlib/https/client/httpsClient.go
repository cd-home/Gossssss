package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ClientWithOutTSL() {
	// 设置不再校验服务端的HTTPS证书
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}

	client := &http.Client{Transport: tr}

	response, err := client.Get("https://127.0.0.1:8081/https-test")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	fmt.Println(string(bytes))
}

func ClientWithTSL() {
	CACrt, err := ioutil.ReadFile("CA.crt")
	if err != nil {
		return
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(CACrt)

	tr := &http.Transport{TLSClientConfig: &tls.Config{RootCAs: pool}}
	client := &http.Client{Transport: tr}

	response, err := client.Get("https://localhost:8081/https-test")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes))
}
