package unsafe

import (
	"testing"
	"unsafe"
)

type example2 struct {
	flag bool // 0xc000100020 <- Starting Address
	// byte            // 0xc000100021 <- 1 byte padding
	counter int16 // 0xc000100022 <- 2 byte alignment
	flag2   bool  // 0xc000100024 <- 1 byte alignment
	// byte            // 0xc000100025 <- 1 byte padding
	// byte            // 0xc000100026 <- 1 byte padding
	//	byte            // 0xc000100027 <- 1 byte padding
	pi float32 // 0xc000100028 <- 4 byte alignment
}

type S struct {
	A uint32
	B uint64
	C uint64
	D uint64
	E struct{}
}

func TestAligin(t *testing.T) {
	e := &example2{
		flag:    true,
		counter: 12,
		flag2:   true,
		pi:      3.145,
	}

	t.Log(unsafe.Alignof(e.flag))

	d := unsafe.Alignof(e)
	t.Log(d)
	t.Log(unsafe.Sizeof(e))

	t.Log(unsafe.Alignof(S{}))
	t.Log(unsafe.Sizeof(S{}))
}
