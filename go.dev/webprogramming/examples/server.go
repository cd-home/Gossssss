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
	mux.HandleFunc("/", YourselfServer)
	// 自定义服务
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}
	// 启动服务
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	// 信号退出
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Shutting down server: %s", err.Error())
	}
}
