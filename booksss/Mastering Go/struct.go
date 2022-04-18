package main

import "log"

type Person struct {
	Name   string
	Height int
	Weight int
}

func main() {
	p1 := Person{
		Name:   "LY",
		Height: 160,
		Weight: 55,
	}
	log.Println(p1)
}
