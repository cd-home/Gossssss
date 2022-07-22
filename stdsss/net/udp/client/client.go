package main

import (
	"flag"
	"fmt"
	"net"
)

var who int

func main() {
	// who
	flag.IntVar(&who, "n", 1, "who")
	flag.Parse()
	// 知道别人是谁
	udpAddr, err := net.ResolveUDPAddr("udp", ":8082")
	if err != nil {
		panic(err)
	}
	client, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// client.WriteToUDP()
	n, err := client.Write([]byte(fmt.Sprintf("From Client [%d] Data", who)))
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
	buf := make([]byte, 512)
	n, _, err = client.ReadFromUDP(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf[:n]))

	// 知道别人是谁, 直接发送
	client.Write([]byte("OK2"))
}
