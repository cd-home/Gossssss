package example_test

import (
	"log"
	"testing"
)

// Pointer
// 指针

func incrementByV(inc int) {
	inc++
	log.Printf("incrementByV: inc: %v, addr: %p", inc, &inc)
}

func incrementByP(inc *int) {
	*inc++
	log.Printf("incrementByP: inc: %v, addr: %p", *inc, inc)
}

func TestMoveByValue(t *testing.T) {
	counter := 10
	log.Printf("TestMoveByValue1: counter: %v, addr: %p", counter, &counter)

	incrementByV(counter)
	log.Printf("TestMoveByValue2: counter: %v, addr: %p", counter, &counter)

	incrementByP(&counter)
	log.Printf("TestMoveByValue3: counter: %v, addr: %p", counter, &counter)
}
