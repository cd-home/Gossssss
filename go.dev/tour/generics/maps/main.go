package main

import (
	"fmt"
	"golang.org/x/exp/maps"
)

func main() {
	m := make(map[int]string, 10)
	m[1] = "1"
	m[2] = "2"
	m[3] = "3"
	fmt.Println(maps.Keys(m))
	fmt.Println(maps.Values(m))
}
