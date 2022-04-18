package main

import "flag"

func main() {
	k := flag.Bool("K", true, "K")
	o := flag.Int("O", 1, "O")
	flag.Parse()

	println(*k)
	println(*o)
}
