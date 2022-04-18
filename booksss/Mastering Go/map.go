package main

import "log"

func main() {
	//m := map[string]string{}
	//m = nil
	var m map[string]string

	log.Println(m)
	m["key"] = "value"
}
