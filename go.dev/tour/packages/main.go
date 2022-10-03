// main package
// Every Go program is made up[组织] of packages.
// Programs start running in package `main`.
package main

// import package by path
import (
	"fmt"
	"math/rand"
	"time"
)

// "入口函数"
func main() {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(100)
	fmt.Println("My favorite number is", randNum)
}
