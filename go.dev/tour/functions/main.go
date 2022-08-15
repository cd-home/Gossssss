package main

import "fmt"

// Foo name
// params list
// return list [support multi return]
func Foo(param1 string, param2 string) (string, error) {
	return param1 + param2, nil
}

func main() {
	// Call Func
	s, e := Foo("a", "b")
	if e != nil {
		panic(e)
	}
	fmt.Println(s)
}
