package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, e := rpc.Dial("tcpdemo", ":8081")
	if e != nil {
		panic(e)
	}
	defer client.Close()
	var reply string
	client.Call("Greeter.Hello", "GodYao", &reply)
	fmt.Println(reply)
}
