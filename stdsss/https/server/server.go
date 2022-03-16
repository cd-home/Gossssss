package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/https-test", Hello)
	// 服务端 数字证书 与 私钥
	err := http.ListenAndServeTLS(":8081", "cert.pem", "key.pem", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "Hello Https!")
}
