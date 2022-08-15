// main packages
package main

// import packages [stdlib]
import (
	"Gossssss/go.dev/tour/packages/exports"
	"fmt"
	"math/rand"
	"time"
)

func ExportNameSpace() {
	exps := exports.ExportStruct{Name: "exports"}
	fmt.Println(exps)

	fmt.Println(exports.ExportConst)
	fmt.Println(exports.ExportVar)
}

func main() {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(100)
	fmt.Println("My favorite number is", randNum)

	ExportNameSpace()

	getter := exports.ExportFunc()
	fmt.Println(getter)
}
