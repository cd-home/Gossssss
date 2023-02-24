package main

import (
	"fmt"
	"net/http"
)

func GetRequest(writer http.ResponseWriter, request *http.Request) {
	// Form-Data Body Fields
	//request.ParseForm()
	//request.PostForm.Get()

	// Not include URL (POST)
	//request.PostFormValue()

	// URL AND Form-Data Body
	//request.ParseForm()
	//request.Form.Get()

	// Use this OK !
	//request.FormValue()
	val := request.FormValue("name")
	fmt.Println(val)

	// One
	//request.ParseMultipartForm()
	//request.MultipartForm.File[""]

	// Two
	//f, header, _ := request.FormFile("upload")
}

func main() {
	http.HandleFunc("/params", GetRequest)
	_ = http.ListenAndServe(":8080", nil)
}
