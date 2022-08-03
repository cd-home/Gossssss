package fmt_test

import (
	"fmt"
	"testing"
)

func TestPaddingZero(t *testing.T) {
	fmt.Println("Hello, world!")

	// %f == float
	i := 12.03
	fmt.Printf("%.4f\n", i)

	// %v == "value"
	fmt.Printf("")

	// %t == type

	// %p == pointer

	// %s == string

	// %d == int
}
