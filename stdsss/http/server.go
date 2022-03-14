package http

import (
	"net/http"
	"time"
	"log"
	"fmt"
)

func YourselfServer(rw http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(rw, "YourselfServer!, Req: %s", r.URL)
}

func DefineYourServer() {

	mux := http.DefaultServeMux
	// mux := http.NewServeMux()

	mux.HandleFunc("/", YourselfServer)

	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
		ReadTimeout: 3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}