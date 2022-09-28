package main

import (
	"fmt"
	"reflect"
)

func main() {
	var m = map[string]string{"a": "b"}
	// Type
	rt := reflect.TypeOf(m)

	if rt.Kind() == reflect.Map {
		fmt.Println("Kind Type Is Map")
	}
}
