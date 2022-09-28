package main

import (
	"fmt"
	"reflect"
)

func main() {
	var s = []int{1, 2, 3, 4}
	rt := reflect.TypeOf(s)

	rv := reflect.ValueOf(s)
	if rt.Kind() == reflect.Slice || rt.Kind() == reflect.Array {
		// for
		for i := 0; i < rv.Len(); i++ {
			fmt.Println(rv.Index(i))
			fmt.Println()
		}
	}
	// set
	rv.Index(0).SetInt(10)
	fmt.Println(s)
}
