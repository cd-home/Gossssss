package patterns_test

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestPipeline(t *testing.T) {
	generator := func(done <-chan struct{}, integers ...int) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for _, integer := range integers {
				select {
				case <-done:
					return
				case intStream <- integer:
				}
			}
		}()
		return intStream
	}

	multiply := func(done <-chan struct{}, intStream <-chan int, multiplier int) <-chan int {
		multipliedStream := make(chan int)
		go func() {
			defer close(multipliedStream)
			for i := range intStream {
				select {
				case <-done:
					return
				case multipliedStream <- i * multiplier:
				}
			}
		}()
		return multipliedStream
	}

	add := func(done <-chan struct{}, initStream <-chan int, additive int) <-chan int {
		addedStream := make(chan int)
		go func() {
			defer close(addedStream)
			for i := range initStream {
				select {
				case <-done:
					return
				case addedStream <- i + additive:
				}
			}
		}()
		return addedStream
	}
	done := make(chan struct{})
	defer close(done)
	intStream := generator(done, 1, 2, 3, 4, 5)
	for i := range multiply(done, add(done, multiply(done, intStream, 2), 1), 2) {
		fmt.Println(i)
	}
}

func TestPipeline2(t *testing.T) {
	repeat := func(done <-chan struct{}, values ...interface{}) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
			}
		}()
		return valueStream
	}

	take := func(done <-chan struct{}, valueStream <-chan interface{}, num int) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}
	done := make(chan struct{})
	defer close(done)
	for i := range take(done, repeat(done, 1), 10) {
		fmt.Printf("%v ", i)
	}
	fmt.Println()
	repeatFn := func(done <-chan struct{}, fn func() interface{}) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- fn():
				}
			}
		}()
		return valueStream
	}
	done2 := make(chan struct{})
	defer close(done2)
	rand := func() interface{} { return rand.Int() }
	for num := range take(done2, repeatFn(done2, rand), 10) {
		fmt.Println(num)
	}
}
