package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// 第三方的 routing 库
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/{user}/{name}/", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Println(vars["user"], vars["name"])
	})

	// assets 文件服务
	fs := http.FileServer(http.Dir("static/"))
	// Need user New router to handle
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	_ = http.ListenAndServe(":8080", router)
}
