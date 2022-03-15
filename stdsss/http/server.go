package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func YourselfServer(rw http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(rw, "YourselfServer!, Req: %s", r.URL)
}

func DefineYourServer() {

	mux := http.DefaultServeMux
	// mux := http.NewServeMux()

	mux.HandleFunc("/", YourselfServer)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancle := context.WithTimeout(context.Background(), time.Second*5)
	defer cancle()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Shutting down server: %s", err.Error())
	}
}
