/*
 * @Time    : 2020年11月22日 17:25:23
 * @Author  : apple-liyao
 * @Project : GodExample
 * @File    : main.go
 * @Software: GoLand
 * @Describe:
 */
package main

import (
	"log"
	"plugin"
)

func main()  {
	p, err := plugin.Open("./tests/plugin.o")
	if err != nil {
		log.Println(err)
	}
	object, err := p.Lookup("FuckFunc")
	if err != nil {
		log.Println(err)
	}
	FuckFunc, ok := object.(func())
	if !ok {
		log.Println("No this func")
	}
	FuckFunc()
}