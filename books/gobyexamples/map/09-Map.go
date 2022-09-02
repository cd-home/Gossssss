package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2

	fmt.Println(m)

	fmt.Println(len(m))


	delete(m, "a")

	fmt.Println(m)

	if v, ok := m["a"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("No a")
	}

}
