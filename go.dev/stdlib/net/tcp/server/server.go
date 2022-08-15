package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	// ResolveTCPAddr 主要是验证协议、地址、端口是否正确
	tcpAddr, err := net.ResolveTCPAddr("tcpdemo", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	tcpListener, err := net.ListenTCP("tcpdemo", tcpAddr)

	if err != nil {
		panic(err)
	}
	for {
		conn, err := tcpListener.AcceptTCP()
		if err != nil {
			continue
		}
		// conn.SetKeepAlive(true)
		go handleConnect(conn)
	}
}

func handleConnect(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}(conn)
	for {
		// 读固定大小, binary 方便数字与字节之间转换, 根据 数字 类型 大小 去读字节
		var dataLen uint32
		binary.Read(conn, binary.LittleEndian, &dataLen)

		fmt.Println("dataLen:", dataLen)

		// 具体非固定大小数据, 需要指定长度读取
		data := make([]byte, dataLen)
		io.ReadFull(conn, data)

		str := string(data)
		fmt.Println("data:", str)
		if str == "quit" {
			break
		}
	}
}

/*
func handleConnect2(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		var dataLen uint32
		binary.Read(reader, binary.LittleEndian, &dataLen)

		fmt.Println("dataLen:", dataLen)
		data := make([]byte, dataLen)
		_, err := reader.Read(data)
		if err != nil {
			continue
		}
		str := string(data)
		fmt.Println("data: ", str)
		if str == "quit" {
			break
		}
	}
}
*/
