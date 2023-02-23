package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, e := rpc.Dial("tcp", ":8081")
	if e != nil {
		panic(e)
	}
	defer func(client *rpc.Client) {
		err := client.Close()
		if err != nil {
			log.Println(err)
		}
	}(client)
	var reply string
	err := client.Call("Greeter.Hello", "GodYao", &reply)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(reply)
}
