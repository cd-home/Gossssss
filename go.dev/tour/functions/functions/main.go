package main

import "fmt"

// Foo function name
// Params list
// Return list [support multi return Values]
// 基本的函数模式
func Foo(param1 string, param2 string, x, y int) (string, int, error) {
	return param1 + param2, x + y, nil
}

// SameParamsType
// When two or more consecutive named function parameters share a type[分享一个类型],
// You can omit[省略] the type from all but the last.
// 多个参数共享相同类型, 可以省略前面的类型, 除了最后一个
func SameParamsType(x, y int) int {
	return x + y
}

// MultiReturnValue
// Return list [support multi return]
// A function can return any number of results.
func MultiReturnValue(x, y int) (int, int) {
	return y, x
}

// NamedReturnValue
// Func return values can be named
// They are treated as variables defined at the top of the function.
func NamedReturnValue(x, y int) (sum int) {
	sum = x + y
	// return can omit sum, if the func is short
	// if this is a long func, better to return values
	return
}

// VariableLengthParameter
// 可变长的参数, 可参考内置函数 append()
func VariableLengthParameter(start int, data ...int) {
	for _, d := range data {
		start += d
	}
	fmt.Println(start)
}

func main() {
	// Call Func
	str, num, e := Foo("a", "b", 1, 2)
	if e != nil {
		panic(e)
	}
	fmt.Println(str)
	fmt.Println(num)

	sums := SameParamsType(1, 2)
	fmt.Println(sums)

	y, x := MultiReturnValue(2, 5)
	fmt.Println(x)
	fmt.Println(y)

	sums = NamedReturnValue(4, 9)
	fmt.Println(sums)

	data := []int{2, 3, 4}
	VariableLengthParameter(2, data...)
}
