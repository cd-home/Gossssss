package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// seed
	rand.Seed(time.Now().Unix())

	fmt.Println(rand.Intn(100))

	fmt.Println(rand.Float64())
	fmt.Println((rand.Float64() * 5) + 5)

	// or 直接用seed更加简单
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	fmt.Println(r.Intn(100))
}
