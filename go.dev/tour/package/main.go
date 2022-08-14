// main package
package main

// import package [std]
import (
	"fmt"
	"math/rand"
	"time"
)

// 
func main() {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(100)
	fmt.Println("My favorite number is", randNum)
}
