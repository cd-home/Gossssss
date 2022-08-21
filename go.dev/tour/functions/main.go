package main

import "fmt"

// Foo name
// Params list
// Return list [support multi return Values]
func Foo(param1 string, param2 string, x, y int) (string, int, error) {
	return param1 + param2, x + y, nil
}

// SameParamsType
// When two or more consecutive named function parameters share a type,
// You can omit[省略] the type from all but the last.
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
}
