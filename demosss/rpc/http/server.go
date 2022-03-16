package main

import (
	"log"
	"net/http"
	"net/rpc"
)

type Greeter struct {}
// 严格符合这种格式的方法，就可以对外提供服务
func (g *Greeter) Hello(req string, resp *string) error {
	*resp = "hello, " + req
	return nil
}


func main() {
	// 注册RPC服务
	_ = rpc.Register(new(Greeter))

	// 提供HTTP形式对外服务
	rpc.HandleHTTP()
	log.Fatal(http.ListenAndServe(":8081", nil))
}
