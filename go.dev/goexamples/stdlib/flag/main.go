package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	strings := os.Args[1:]
	fmt.Println(strings)


	w := flag.String("word", "foo", "a string")
	f := flag.Bool("fork", false, "a bool")
	n := flag.Int("num", 100, "a number")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a var string")

	flag.Parse()

	fmt.Println("word", *w)
	fmt.Println("fork", *f)
	fmt.Println("num", *n)

	fmt.Println("svar", svar)
}
