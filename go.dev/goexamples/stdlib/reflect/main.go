package main

import (
	"fmt"
	"os"
	"reflect"
)

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

func ReflectStructType() {
	a := A{
		X: 100,
		Y: 200,
		Z: "A",
	}
	r1 := reflect.ValueOf(&a).Elem()

	for i := 0; i < r1.NumField(); i++ {
		fmt.Print(r1.Type().Field(i).Name, " ")
		fmt.Print(r1.Type().Field(i).Type, " ")
		fmt.Print(r1.Field(i).Interface(), " ")
		fmt.Println()
	}
}

func PrintMethods(i interface{}) {
	r := reflect.ValueOf(i)
	t := r.Type()
	fmt.Printf("%s\n", t)
	for j := 0; j < r.NumMethod(); j++ {
		m := r.Method(j).Type()
		fmt.Println(t.Method(j).Name, "-->", m)
	}
}

type D struct {
	X    int
	Y    int
	Text string
}

func (d *D) CompareStruct(dd D) bool {
	d1 := reflect.ValueOf(d).Elem()
	d2 := reflect.ValueOf(&dd).Elem()

	for i := 0; i < d1.NumField(); i++ {
		if d1.Field(i).Interface() != d2.Field(i).Interface() {
			return false
		}
	}
	return true
}

func main() {
	ReflectStructType()

	var f *os.File
	PrintMethods(f)

	d1 := D{
		X:    1,
		Y:    1,
		Text: "a",
	}
	d2 := D{
		X:    1,
		Y:    1,
		Text: "a",
	}
	fmt.Println(d1.CompareStruct(d2))
}
