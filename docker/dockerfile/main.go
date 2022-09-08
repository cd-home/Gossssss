package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		bs, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}
		fmt.Fprintf(rw, string(bs))
	})
	_ = http.ListenAndServe(":8999", nil)

}
