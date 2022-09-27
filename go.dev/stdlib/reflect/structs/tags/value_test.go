package main

import (
	"reflect"
	"testing"
)

func TestTagValue(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}
	var u = User{}
	rv := reflect.ValueOf(&u)
	f := rv.Elem().FieldByName("Age")
	f.SetInt(10)
	t.Log(u)

	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		if !field.CanSet() {
			continue
		}
	}

	// judge the field IsExported
	rt := reflect.TypeOf(&u)
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		if !field.IsExported() {
			continue
		}
		// Means that the IsExported
		if len(field.PkgPath) == 0 {
			continue
		}
	}

}
