package exercise_test

import (
	"fmt"
	"testing"
)

type ByteCounter int

// Write
// 自定义类型的 Write 方法
func (bc *ByteCounter) Write(p []byte) (n int, err error) {
	*bc += ByteCounter(len(p))
	return len(p), nil
}

func TestContract(t *testing.T) {
	// self call
	var bc ByteCounter
	bc.Write([]byte("Hello"))
	fmt.Println(bc)

	// by interface call
	bc = 0
	name := "Dolly"
	fmt.Fprintf(&bc, "Hello, %s", name)
	fmt.Println(bc)
}
