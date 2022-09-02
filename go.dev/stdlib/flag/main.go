package main

import (
	"flag"
	"fmt"
	"time"
)

func SimpleFlags() {
	// support
	// -from=x  --from=y
	// -from x  --from y
	var from = flag.String("from", "ch", "where from")

	// support true [True] false [False]
	var stop = flag.Bool("stop", false, "stop")

	// support 1h 1h2m
	var dur = flag.Duration("time", time.Hour, "")

	flag.Parse()
	fmt.Println(*from)
	fmt.Println(*stop)
	fmt.Println(*dur)
}

func VarFlags() {
	var ip string
	var port string
	flag.StringVar(&ip, "ip", "127.0.0.1", "IP address")
	flag.StringVar(&port, "port", ":8080", "Port number")
	flag.Parse()
	// ./flag -ip=127.0.0.1 -port=8989
	fmt.Println(ip, port)
}

func main() {
	//SimpleFlags()
	VarFlags()
	//MoreFlags()
}
