/*
 * @Time    : 2020年11月22日 17:23:58
 * @Author  : apple-liyao
 * @Project : GodExample
 * @File    : plugin.go
 * @Software: GoLand
 * @Describe:
 */
package main

import "fmt"

func FuckFunc() {
	fmt.Println("secret fuck func")
}

// go build -buildmode=plugin -o plugin.o plugin.go

