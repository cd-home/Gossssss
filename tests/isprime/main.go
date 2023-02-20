package main

import "fmt"

/*
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
What is the 10 001st prime number?
prime number： 质数
*/

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	count := 0
	num := 2
	for count <= 10001 {
		if isPrime(num) {
			count++
		}
		num++
	}
	fmt.Println(num)
}
