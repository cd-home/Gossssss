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
	user := User{Name: "yao", Age: 28}

	rv := reflect.ValueOf(&user)

	// By MethodName Call
	f := rv.MethodByName("SayHello")
	f.Call([]reflect.Value{})

	// Call Methods By MethodName and With Params
	Invoke(&user, "SayHello")
	Invoke(&user, "SayHi", "Hi")

	// Func IsExported
	rt := reflect.TypeOf(&user)
	e := rt.Elem()
	for i := 0; i < e.NumMethod(); i++ {
		m := e.Method(i)
		fmt.Println(m.Name)
		if m.Name == "SayHello" && m.IsExported() {
			// one way
			rv.MethodByName(m.Name).Call([]reflect.Value{})
			// another way, Func must use receiver be the first argument
			m.Func.Call([]reflect.Value{reflect.ValueOf(user)})
		}
	}

	// Got Fields
	ev := rv.Elem()
	for i := 0; i < ev.NumField(); i++ {
		field := ev.Field(i)
		fmt.Printf("%d %s = %v \n", i, field.Type(), field.Interface())
	}
}

func Invoke(o interface{}, funcName string, params ...interface{}) {
	rv := reflect.ValueOf(o)
	ins := make([]reflect.Value, len(params))
	// Real params
	for i, value := range params {
		ins[i] = reflect.ValueOf(value)
	}
	// Call Method
	f := rv.MethodByName(funcName)
	f.Call(ins)
}
