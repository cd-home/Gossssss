package main

import (
	"fmt"
	"os"
)

func main() {
	//os.Setenv("TEST", "TEST1")
	value := os.Getenv("ENV_TAG")

	fmt.Println(value)

	for k, v := range os.Environ() {
		fmt.Println(k, v)
	}
}
