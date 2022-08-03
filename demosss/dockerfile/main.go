package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		bs, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}
		f, err := os.OpenFile("news.json", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
		if err != nil {
			log.Println(err)
		}

		defer f.Close()
		f.WriteString(string(bs) + "\n")
	})
	_ = http.ListenAndServe(":8999", nil)

}
