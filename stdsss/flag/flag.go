package main

import (
	"flag"
	"fmt"
)


func main() {
	var ip string
	var port string
	flag.StringVar(&ip, "ip", "127.0.0.1", "IP address")
	flag.StringVar(&port, "port", ":8080", "Port number")
	flag.Parse()

	// ./flag -ip=127.0.0.1 -port=8989
	fmt.Println(ip, port)
}
