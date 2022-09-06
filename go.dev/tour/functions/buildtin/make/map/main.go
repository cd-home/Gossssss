package main

import "fmt"

func main() {
	// map
	m := make(map[string]string)
	m["name"] = "yao"
	fmt.Println(m)
	fmt.Println(len(m))

	//make(map, cap)
	mCap := make(map[string]string, 10)
	mCap["name"] = "yao"
	fmt.Println(mCap)
	fmt.Println(len(mCap))
}
