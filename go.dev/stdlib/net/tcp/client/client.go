package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcpdemo", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	tcpConn, err := net.DialTCP("tcpdemo", nil, tcpAddr)
	if err != nil {
		panic(err)
	}
	defer tcpConn.Close()
	// tcpConn.SetKeepAlive(true)
	buf := new(bytes.Buffer)
	for i := 1; i < 6; i++ {
		str := ""
		for j := 0; j < i; j++ {
			str += "Hello, world! "
		}
		dataLen := len(str)
		fmt.Println(dataLen)
		binary.Write(buf, binary.LittleEndian, uint32(dataLen))
		fmt.Println(str)
		binary.Write(buf, binary.LittleEndian, []byte(str))
	}
	quit := "quit"
	binary.Write(buf, binary.LittleEndian, uint32(len(quit)))
	binary.Write(buf, binary.LittleEndian, []byte(quit))
	tcpConn.Write(buf.Bytes())
}
