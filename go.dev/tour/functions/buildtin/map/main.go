package main

import "fmt"

func main() {
	// map
	m := make(map[string]string)
	m["name"] = "yao"
	m["age"] = "27"
	fmt.Println(m)

	delete(m, "age")
	fmt.Println(m)
}
