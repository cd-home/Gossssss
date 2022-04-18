package main

import "log"

const (
	B = 1 << (iota * 10)
	KB
	MB
	GB
)

const (
	a = iota
	_
	c
	_
	e
)

func main() {
	log.Println(B)
	log.Println(KB)
	log.Println(MB)
	log.Println(GB)

	log.Println(1 << 0)
	log.Println(1 << 10)
	log.Println(1 << 20)

	log.Println(a)
	log.Println(c)
	log.Println(e)
}
