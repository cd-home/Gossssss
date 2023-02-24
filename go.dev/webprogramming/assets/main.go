package main

import (
	"net/http"
)

func main() {
	// 静态文件服务
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	_ = http.ListenAndServe(":8080", nil)
}
