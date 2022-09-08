package main

import (
	"Gossssss/docker/docker-compose/app/cache"
	"bytes"
	"context"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", HelloWorld)
	_ = http.ListenAndServe(":8080", nil)
}

func HelloWorld(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers", "*")
	cache.Rdb.Incr(context.Background(), "app_count")
	count, _ := cache.Rdb.Get(context.Background(), "app_count").Result()
	env := os.Getenv("ENV")
	result := count + env
	data := bytes.NewReader([]byte(result))
	_, _ = io.Copy(rw, data)
}
