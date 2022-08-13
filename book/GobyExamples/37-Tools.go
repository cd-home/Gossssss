package main

import (
	"fmt"
	"strings"
)

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func Filter(vs []string, f func(string) bool) []string {
	var vsf []string
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {
	fruits := []string{"apple", "banana", "peach", "kiwi"}

	fmt.Println(Index(fruits, "banana"))

	fmt.Println(Include(fruits, "abc"))

	fmt.Println(Any(fruits, func(s string) bool {
		return strings.HasPrefix(s, "p")
	}))

	fmt.Println(Filter(fruits, func(s string) bool {
		return strings.Contains(s, "a")
	}))

	fmt.Println(Map(fruits, func(s string) string {
		return strings.ToUpper(s)
	}))
}
