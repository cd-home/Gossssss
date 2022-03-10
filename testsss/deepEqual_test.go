package testsss


import (
	"testing"
	"reflect"
)

type TestSuite struct {
	a, b, c string
}

func TestDeepEquals(t *testing.T) {
	m := &TestSuite{"a", "b", "c"}
	n := &TestSuite{"a", "b", "c"}

	// 显示转 interface{}
	x := interface{}(m)
	y := interface{}(n)

	t.Log(m == n) // false
	t.Log(x == y) // false
	// 不转参数默认转
	t.Log(reflect.DeepEqual(m, n)) // true
	t.Log(reflect.DeepEqual(x, y)) // true
}