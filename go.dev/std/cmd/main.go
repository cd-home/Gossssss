package main

import (
	"Gossssss/go.dev/std/cmd/build"
	"fmt"
)

func main() {
	// go build -gcflags '-m -l' main.go
	value := build.Escape(1)
	fmt.Println(value, *value)
}
