package main

import (
	"net/http"
	api "Gossssss/stdsss/http"
)

func main() {
	http.HandleFunc("/", api.HelloWorld)

	api.FileServer()

	http.ListenAndServe(":8080", nil)
}