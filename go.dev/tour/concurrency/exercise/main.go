package main

import (
	"Gossssss/go.dev/tour/concurrency/exercise/tree"
	"fmt"
)

func main() {
	t1 := tree.New(10)
	t2 := tree.New(11)
	fmt.Println(t1.String())
	fmt.Println(t2.String())
	fmt.Println(tree.Same(t1, t2))
}
