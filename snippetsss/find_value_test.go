package snippetsss

import (
	"reflect"
	"testing"
)

func FindValueFromSlice(value interface{}, container interface{}) (bool, int) {
	exist := false 
	idx := -1
	if reflect.TypeOf(container).Kind() == reflect.Slice {
		cv := reflect.ValueOf(container)
		for idx = 0; idx < cv.Len(); idx++ {
			if reflect.DeepEqual(value, cv.Index(idx).Interface()) {
				return !exist, idx
			}
		}
	}
	return exist, idx
}

func TestFindValueFromSlice(t *testing.T) {
	value := 2
	slices := []int{1, 2, 3, 4, 5}
	t.Log(FindValueFromSlice(value, slices))
}
