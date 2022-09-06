package x

import (
	"testing"
	"unsafe"
)

func TestCompareEmptyStruct(t *testing.T) {
	s1 := new(struct{})
	s2 := new(struct{})

	// 逃逸
	t.Logf("%p", s1)
	t.Logf("%p", s2)

	t.Log(unsafe.Sizeof(s1))
	t.Log(unsafe.Sizeof(s2))

	t.Log(s1 == s2)

	t.Log("=========================================")

	// 没有逃逸
	s3 := new(struct{})
	s4 := new(struct{})

	t.Log(unsafe.Sizeof(s3))
	t.Log(unsafe.Sizeof(s4))
	t.Log(s3 == s4)
}
