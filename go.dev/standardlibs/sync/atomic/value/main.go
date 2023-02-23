package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	config := make(map[string]string)
	config["name"] = "yao"

	var anyValue atomic.Value
	anyValue.Store(config)

	val := anyValue.Load()
	c, ok := val.(map[string]string)
	if ok {
		fmt.Println(c["name"])
	}
}
