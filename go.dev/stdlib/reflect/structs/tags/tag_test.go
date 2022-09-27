package main

import (
	"reflect"
	"testing"
)

type Result struct {
	Name string `http:"name_me"`
	Age  string `http:"age_me"`
}

func TestTag(t *testing.T) {
	// mock request body data
	data := map[string]string{
		"name_me": "Mike",
		"age_me":  "12",
	}

	// result data
	var ret Result

	fields := make(map[string]reflect.Value)

	// struct
	rv := reflect.ValueOf(&ret).Elem()

	for i := 0; i < rv.NumField(); i++ {
		fieldInfo := rv.Type().Field(i)
		tag := fieldInfo.Tag.Get("http")
		t.Log(tag)
		fields[tag] = rv.Field(i)
	}

	t.Log(fields)

	for k, v := range data {
		t.Log(k)
		f := fields[k]
		if !f.IsValid() {
			continue
		}
		if f.Kind() == reflect.String {
			f.SetString(v)
		}
	}
	t.Log(ret)
}
