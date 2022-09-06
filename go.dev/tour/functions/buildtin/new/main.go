package main

import "fmt"

func main() {
	// new: return pointer
	type User struct {
		Name string
	}
	user := new(User)
	user.Name = "yao"
	fmt.Println(user)

	var p = new(int)
	*p = 19
	fmt.Println(*p)
}
