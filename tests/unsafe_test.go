package testsss

import (
	"testing"
	"unsafe"
	"reflect"
)

func TestUnsafePointer(t *testing.T) {
	bs := []byte("123")

	// str pointer to bs
	str := *(*string)(unsafe.Pointer(&bs))

	bs[1] = '4'

	t.Logf("%+v", string(bs))
	t.Logf("%+v", string(str))
}

func TestUnsafePointerOriginSlice(t *testing.T) {
	sli := make([]int, 10)

	for i := range sli {
		sli[i] = i
	}

	sub := sli[:5]

	t.Logf("sli: %p, %+v\n", &sli, sli) 
	t.Logf("sub: %p, %+v\n", &sub, sub)


	// donot know this meening
	data := (*[5]int)(unsafe.Pointer(((*reflect.SliceHeader)(unsafe.Pointer(&sub))).Data))
	t.Logf("data: original\t%p:%+v\n", data, *data)

}
