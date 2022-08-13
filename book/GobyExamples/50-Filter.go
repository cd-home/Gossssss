package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// cat ./46-URL.go | go run 50.Filter.go
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}
	if err := scanner.Err(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "error", err)
		os.Exit(1)
	}
}
