package main

import (
	"fmt"
	"net/rpc"
	"testing"
)

func TestRpcClient(t *testing.T) {
	client, _ := rpc.Dial("tcp", "127.0.0.1:8081")

	var resp string
	_ = client.Call("Greeter.Hello", "GodYao", &resp)

	fmt.Println(resp)
}
