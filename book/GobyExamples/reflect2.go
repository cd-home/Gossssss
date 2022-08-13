package main

import (
	"fmt"
	"os"
	"reflect"
)

type t1 int

type t2 int

type D struct {
	X int
	Y int
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

func PrintMethods(i interface{}) {
	r := reflect.ValueOf(i)
	t := r.Type()
	fmt.Printf("%s\n", t)
	for j := 0; j < r.NumMethod(); j++ {
		m := r.Method(j).Type()
		fmt.Println(t.Method(j).Name, "-->", m)
	}
}

func main() {
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