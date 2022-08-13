package example_test

import (
	"fmt"
	"testing"
)

func TestTypeZeroValue(t *testing.T) {
	var a int
	var b string
	var c float64
	var d bool
	fmt.Printf("var a int \t type: %T     value: %v\n", a, a)
	fmt.Printf("var b string \t type: %T  value: %v\n", b, b)
	fmt.Printf("var c float64 \t type: %T value: %v\n", c, c)
	fmt.Printf("var d bool \t type: %T    value: %v\n\n", d, d)
}

func TestShortDeclareType(t *testing.T) {
	aa := 10      // int 		[10]
	bb := "hello" // string 	[hello]
	cc := 3.14159 // float64 	[3. 14159]
	dd := true    // bool 		[true]
	fmt.Printf("aa := 10 \t type: %T     	value:[%v]\n", aa, aa)
	fmt.Printf("bb := \"hello\" \t type: %T   value:[%v]\n", bb, bb)
	fmt.Printf("cc := 3. 14159 \t type: %T  value:[%v]\n", cc, cc)
	fmt.Printf("dd := true \t type: %T     value:[%v]\n\n", dd, dd)
}
