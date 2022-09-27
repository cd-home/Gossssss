package main

import (
	"fmt"
	"reflect"
)

func main() {
	m := map[string]string{"one": "1", "two": "2"}
	vm := reflect.ValueOf(m)
	if vm.Kind() == reflect.Map {
		for _, key := range vm.MapKeys() {
			fmt.Println(vm.MapIndex(key))
		}
	}
	tf2 := reflect.TypeOf(m)
	fmt.Println(tf2.Size())
}
