package main

import "C"
import "log"

//export GoFunc
func GoFunc()  {
	log.Println("Go func")
}

func main() {

}
