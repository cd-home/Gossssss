package main

import (
	"Gossssss/third_party/grpc/lb/etcdv3/register"
	"flag"
	"fmt"
	"log"
	"strings"
)

var endpoints string
var service string
var lease int64
var host string
var port string

func init() {
	flag.StringVar(&endpoints, "endpoints", "127.0.0.1:2379", "endpoints")
	flag.StringVar(&host, "host", "127.0.0.1", "host")
	flag.StringVar(&port, "port", "8080", "port")
	flag.StringVar(&service, "service", "hello", "service")
	flag.Int64Var(&lease, "lease", 5, "lease")
}

func main() {
	_endpoints := strings.Split(endpoints, ",")
	addr := fmt.Sprintf("%s:%s", host, port)
	sr, err := register.New(_endpoints, service, addr, lease)
	if err != nil {
		log.Fatal(err)
	}
	defer sr.Revoke()
	select {}
}
