package main

import "fmt"

func main() {
	// make return the "type" self
	// slice
	s := make([]int, 10, 20)
	fmt.Println(s)

	// map
	m := make(map[string]string, 10)
	m["name"] = "yao"
	fmt.Println(m)

	// channel
	ch := make(chan string)
	go func() {
		ch <- "yao"
	}()
	fmt.Println(<-ch)

	// new: return pointer
	type User struct {
		Name string
	}
	user := new(User)
	user.Name = "yao"
	fmt.Println(user)
}
