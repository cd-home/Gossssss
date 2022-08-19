package example_test

import (
	"fmt"
	"testing"
)

func TestConvType(t *testing.T) {
	var a int8 = 10
	b := int32(a)
	fmt.Printf("a := 10 type: %T value: %v\n", a, a)
	fmt.Printf("b := int32(10) type: %T value: %v\n", b, b)
}
