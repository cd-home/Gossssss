package main

import (
	api "Gossssss/stdsss/http"
	"net/http"
)

func main() {
	http.HandleFunc("/", api.HelloWorld)

	http.HandleFunc("/upload", api.UploadFile)

	api.FileServer()

	http.ListenAndServe(":8080", nil)
}
