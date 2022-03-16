package main

import (
	"net"
	"net/rpc"
)

type Greeter struct {}
// 严格符合这种格式的方法，就可以对外提供服务
func (g *Greeter) Hello(req string, resp *string) error {
	*resp = "hello, " + req
	return nil
}

func main() {
	// 注册
	_ = rpc.Register(new(Greeter))

	// 监听地址，提供服务
	addr, _ := net.ResolveTCPAddr("tcp", ":8081")
	listener, _ := net.ListenTCP("tcp", addr)

	for {
		conn, _ := listener.Accept()
		go rpc.ServeConn(conn)
	}
}
