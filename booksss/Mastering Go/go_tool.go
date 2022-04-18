package main

import (
	"fmt"
	"log"
)

type myint int

type fuckint = int

func main() {
	var i myint = 10
	log.Println(i)
	fmt.Printf("%T", i)
	var x fuckint = 10
	fmt.Printf("%T", x)

}
