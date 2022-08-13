package reflect

import (
	"reflect"
	"testing"
)

func TestReflectKind(t *testing.T) {
	var m = map[string]string{"a": "b"}
	rtk := reflect.TypeOf(m).Kind()
	if rtk == reflect.Map {
		t.Log("Kind Type Is Map")
	}
}
