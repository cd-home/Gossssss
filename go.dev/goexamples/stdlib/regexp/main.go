package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Simple
	match, _ := regexp.MatchString("p[a-z]+ch", "peach")
	fmt.Println(match)

	// Compile
	r := regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peach"))

	fmt.Println(r.FindString("peach punch"))
	fmt.Println(r.FindStringIndex("peach punch"))

	fmt.Println(r.FindStringSubmatch("peach punch"))
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
}
