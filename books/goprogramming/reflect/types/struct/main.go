package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
	sex  string
}

func (u User) GetName() string {
	return u.Name
}

func main() {

	user := User{
		Name: "yao",
		Age:  27,
		sex:  "man",
	}
	rv := reflect.ValueOf(user)
	// 用来获取成员类型信息
	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		// 获取字段名称
		fmt.Println(rt.Field(i).Name)
		// 获取tag
		//fmt.Println(rt.Field(i).Tag.Get("json"))
		// 通过 FieldName获取字段
		//fmt.Println(rt.FieldByName("Name"))
		// 获取值
		fmt.Println(rv.Field(i))
		// 是否可导出
		fmt.Println(rt.Field(i).IsExported())
	}
	// 获取类型的方法数量
	fmt.Println("methods: ", rv.NumMethod())

	fmt.Println("call method:", rv.MethodByName("GetName").Call([]reflect.Value{}))

	// 如果要设置, 必须传递指针
	rvp := reflect.ValueOf(&user)
	// 并且通过 Elem 获取原值
	elem := rvp.Elem()
	elem.FieldByName("Age").SetUint(28)
	fmt.Println(user)
	fmt.Println("methods: ", rvp.NumMethod())
}
