package main

import (
	"errors"
	"fmt"
)

func main() {
	e := errors.New("error is interface")
	fmt.Println(e.Error())
}
