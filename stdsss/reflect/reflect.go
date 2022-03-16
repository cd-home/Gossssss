package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  uint8
}

func (u User) SayHello() {
	fmt.Printf("Hello, I am %s\n", u.Name)
}

func (u User) SayHi(msg string) {
	fmt.Printf("%s, I am %s\n", msg, u.Name)
}

func main() {
	user := User{Name: "GodYao", Age: 27}

	object := reflect.ValueOf(&user)
	// By MethodName Call
	f := object.MethodByName("SayHello")
	f.Call([]reflect.Value{})

	// Got Fields
	t := object.Elem()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("%d %s = %v \n", i, field.Type(), field.Interface())
	}

	//
	Invoke(&user, "SayHello")
	Invoke(&user, "SayHi", "Hi")
}

func Invoke(o interface{}, funcName string, params ...interface{}) {
	object := reflect.ValueOf(o)
	ins := make([]reflect.Value, len(params))
	// Real params
	for i, value := range params {
		ins[i] = reflect.ValueOf(value)
	}
	// Call Method
	f := object.MethodByName(funcName)
	f.Call(ins)
}
