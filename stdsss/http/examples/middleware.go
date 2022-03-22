package http

import (
	"log"
	"net/http"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Host, r.Method, r.URL)
		f(w, r)
	}
}
