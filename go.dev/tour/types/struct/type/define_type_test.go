package _type

import (
	"testing"
)

// define new type based on base type
type Myint int

// alias to base type
type Aliaint = int

func TestDefineMyType(t *testing.T) {
	var num1 int = 123
	// var num2 Myint = 123
	var num3 Aliaint = 123

	// t.Log(num1 == num2)
	// t.Log(num2 == num3)

	t.Log(num1 == num3)
}
