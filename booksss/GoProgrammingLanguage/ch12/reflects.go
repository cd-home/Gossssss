package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main1() {
	//	注意，反射是依赖接口的

	// reflect.Type 表示Go的类型   reflect.Value 装载任意类型的值

	// 获取类型
	// 27是一个具体的值，将其隐式转换为了interface{}, 结果式创建一个具有两个信息的接口值：动态类型 与 动态值
	// 返回一个 reflect.Type
	t := reflect.TypeOf(27)
	fmt.Println(t.String())

	fmt.Printf("%T\n", 27)

	// 返回装载动态值的 reflect.Value
	v := reflect.ValueOf(18)
	fmt.Println(v)
	fmt.Println(v.String())
	// Value 调用Type() 返回动态值的类型
	vt := v.Type()
	fmt.Println(vt.String())

	// Value 可以逆转为 interface{}, 显然Value 比interface{} 所提供的方法更多
	x := v.Interface()
	i := x.(int)
	fmt.Printf("%d\n", i)
}

func Sprint(i interface{}) string {
	type stringer interface {
		String() string
	}
	// 根据不同的输入，做类型选择
	// 只能不断添加case分支, 并且如果有很多的自定义类型，显然这样的方式并不合理
	switch t := i.(type) {
	case stringer:
		return t.String()
	case string:
		return t
	case int:
		return strconv.Itoa(t)
	case bool:
		if t {
			return "true"
		}
		return "false"
	default:
		return "Don not Know Type"
	}
}
