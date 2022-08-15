package main

import (
	"log"
	"net"
	"net/rpc"
)

// RPC: Remote Procedure Call
// Using TCP/IP

type Greeter struct{}

// 严格符合这种格式的方法，就可以对外提供服务
func (g *Greeter) Hello(req string, resp *string) error {
	*resp = "Hello, " + req
	return nil
}

func main() {
	// 注册
	_ = rpc.Register(&Greeter{})

	// 监听地址，提供服务
	addr, _ := net.ResolveTCPAddr("tcpdemo", ":8081")
	listener, _ := net.ListenTCP("tcpdemo", addr)

	for {
		conn, e := listener.AcceptTCP()
		if e != nil {
			continue
		}
		log.Printf("%s\n", conn.RemoteAddr())
		go rpc.ServeConn(conn)
	}
}
