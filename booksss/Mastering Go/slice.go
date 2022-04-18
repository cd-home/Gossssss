package main

import (
	"log"
	"sort"
)

func main() {
	s := make([]int, 20)
	s = append(s, 1)
	log.Println(s)
	s = append(s, []int{1, 2, 3}...)
	log.Println(s)

	a := []int{1, 2, 3, 4, 5, 6}
	b := []int{7, 8, 9, 10}
	copy(a, b)
	log.Println(a)
	log.Println(b)

	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})

	log.Println(s)
}
