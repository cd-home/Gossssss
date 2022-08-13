package uber_test

import "testing"

// Abstract
type F interface {
	foo(string)
}

// Impl
type S1 struct {
	V string
}

func (s S1) foo(v string) {
	s.V = v
}

type S2 struct {
	V string
}

func (s *S2) foo(v string) {
	s.V = v
}

func TestInterfaceMethodModifyOrgValue(t *testing.T) {
	// 注意如果需要修改源值: 1. 需要指针接收者 2. 对象指针赋值接口变量

	s1 := S1{V: "123"}
	s2 := &S2{V: "456"}

	// 赋值接口变量
	var f1 F = s1
	var f2 F = s2

	f1.foo("321")
	f2.foo("654")

	t.Log(s1)
	t.Log(s2)
}
