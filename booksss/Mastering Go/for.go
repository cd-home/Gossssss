package main

import "log"

func main() {
	i := 10
	for  {
		if i < 0 {
			break
		}
		log.Println(i)
		i--
	}
}
