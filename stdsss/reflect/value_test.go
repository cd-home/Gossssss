package reflect

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
	vf := reflect.ValueOf(&u)
	f := vf.Elem().FieldByName("Age")
	f.SetInt(10)
	t.Log(u)
}
