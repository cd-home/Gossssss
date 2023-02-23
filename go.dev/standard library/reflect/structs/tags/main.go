package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

func main() {
	user := &User{
		Name: "yao",
		Age:  28,
	}
	// Type
	fmt.Println("Type........")
	rt := reflect.TypeOf(user).Elem()
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		fmt.Println(f.Name)
		fmt.Println(f.Tag.Get("json"))
		fmt.Println(f.IsExported())
	}

	fmt.Println()
	fmt.Println("Value........")
	// Value
	rv := reflect.ValueOf(user)

	// 由于前面是 &user的指针, 都需要Elem
	velem := rv.Elem()
	telem := rv.Type().Elem()
	if velem.Kind() == reflect.Struct {
		for i := 0; i < velem.NumField(); i++ {
			fmt.Println("CanSet: ", velem.Field(i).CanSet())
			// 字段的类型
			fmt.Println(velem.Field(i).Kind())
			switch velem.Field(i).Kind() {
			case reflect.String:
				// 字段的值
				fmt.Println(velem.Field(i).String())
			case reflect.Uint8:
				fmt.Println(velem.Field(i).Uint())
			}
		}
	}
	// 由于 &User是指针, 所有需要Elem()
	name := velem.FieldByName("Name")
	// 设置值
	name.SetString("mike")
	fmt.Println(name)

	//fmt.Println(rtm.Implements())
	fmt.Println(telem.Name())
}
