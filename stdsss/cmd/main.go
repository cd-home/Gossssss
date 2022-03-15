package main

import (
	build "Gossssss/stdsss/cmd/build"
	"fmt"
)

func main() {
	value := build.Escape(1)
	fmt.Println(value, *value)
}
