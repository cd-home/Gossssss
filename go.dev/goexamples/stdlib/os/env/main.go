package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//os.Setenv("TEST", "TEST1")
	value := os.Getenv("ENV_TAG")

	fmt.Println(value)

	for _, v := range os.Environ() {
		pair := strings.SplitN(v, "=", 2)
		fmt.Println(pair)
	}
}
