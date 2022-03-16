package main

import (
	"fmt"
	"net/rpc"
	"testing"
)

func TestRpcHTTP(t *testing.T) {
	// 服务地址
	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:8081")
	var resp string

	// 调用Greeter服务的Hello方法
	_ = client.Call("Greeter.Hello", "GodYao", &resp)

	fmt.Println(resp)
}
