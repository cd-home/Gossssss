package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println(r.FormValue("id"))
	})
	_ = http.ListenAndServe(":8081", nil)
}
