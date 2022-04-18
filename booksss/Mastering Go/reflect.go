package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := A{
		X: 100,
		Y: 200,
		Z: "A",
	}
	c := C{
		F: 1,
		G: 2,
		I: 1.2,
	}
	r1 := reflect.ValueOf(&a).Elem()

	for i := 0; i < r1.NumField(); i++ {
		fmt.Print(r1.Type().Field(i).Name, " ")
		fmt.Print(r1.Type().Field(i).Type, " ")
		fmt.Print(r1.Field(i).Interface(), " ")
		fmt.Println()
	}

	fmt.Println()

	r2 := reflect.ValueOf(&c).Elem()

	for i := 0; i < r2.NumField(); i++ {
		fmt.Print(r2.Type().Field(i).Name, " ")
		fmt.Print(r2.Type().Field(i).Type, " ")
		fmt.Print(r2.Field(i).Interface(), " ")
		fmt.Println()
	}
}

type A struct {
	X int
	Y int
	Z string
}

type C struct {
	F int
	G int
	I float64
}