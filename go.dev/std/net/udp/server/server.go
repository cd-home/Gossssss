package main

import (
	"fmt"
	"net"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", ":8082")
	if err != nil {
		panic(err)
	}
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		panic(err)
	}
	// udpConn.SetDeadline()
	// defer udpConn.Close()
	buf := make([]byte, 512)
	for {
		// 不知道别人是谁, 所以可以用到from
		n, from, err := udpConn.ReadFromUDP(buf)
		if err != nil {
			panic(err)
		}
		data := buf[:n]
		fmt.Println(string(data))

		n, err = udpConn.WriteToUDP([]byte("OK"), from)
		if err != nil {
			panic(err)
		}
		fmt.Println(n)
	}
}
