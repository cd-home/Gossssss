package main

import (
	"fmt"
	"reflect"
)

func main() {
	m := map[string]string{"name": "yao", "age": "28"}
	rvm := reflect.ValueOf(m)
	if rvm.Kind() == reflect.Map {
		// get
		for _, key := range rvm.MapKeys() {
			fmt.Println(key, rvm.MapIndex(key))
			// set
			rvm.SetMapIndex(reflect.ValueOf("name"), reflect.ValueOf("mike"))
			rvm.SetMapIndex(reflect.ValueOf("age"), reflect.ValueOf("20"))
		}
	}
	// get
	fmt.Println(rvm.MapIndex(reflect.ValueOf("name")))
	fmt.Println(m)

	rtm := reflect.TypeOf(m)
	// key type
	fmt.Println(rtm.Key())
	fmt.Println(rtm.Comparable())
}
