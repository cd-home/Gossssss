package main

import "fmt"

// DeclaresVariable
// Just Declares
var DeclaresVariable string

// SingleGlobalVar
// Package Scope Var
// Declaration And Initializers
var SingleGlobalVar string = "Hello World"

// Type Assets Meantime Diff Type
var a, b, c = 10, 20, true

// Group
var (
	x string = "Go"
	y string = "Python"
	z string = "JavaScript"
)

func main() {
	// Default Zero Value
	fmt.Println(DeclaresVariable)

	fmt.Println(SingleGlobalVar)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// Function Scope Var
	var name string = "Yao"

	fmt.Println(name)

	// Short variable declarations Just In functions
	age := "18"
	fmt.Println(age)
}
