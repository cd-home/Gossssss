package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name"`
	age  uint8  `json:"age"`
}

func main() {
	var i interface{} = &User{
		Name: "yao",
		age:  28,
	}
	tf := reflect.TypeOf(i).Elem()
	for i := 0; i < tf.NumField(); i++ {
		f := tf.Field(i)
		fmt.Println(f.Tag.Get("json"))
		fmt.Println(f.IsExported())
	}

	// need struct
	vf := reflect.ValueOf(i)
	if !vf.IsNil() {
		vf = vf.Elem()
	}
	if vf.Kind() == reflect.Struct {
		fmt.Println(vf.Field(0).CanSet())
		fmt.Println(vf.Field(1).CanSet())
	}
	name := vf.Elem().FieldByName("Name")
	name.SetString("mike")
	fmt.Println(name)
}
