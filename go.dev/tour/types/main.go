package main

import "fmt"

func InitMap(cap uint8) map[string]string {
	dataMap := make(map[string]string, cap)
	return dataMap
}

func NilMapPanic() map[string]string {
	var maps map[string]string
	return maps
}

func main() {
	// make map
	maps := InitMap(10)
	maps["name"] = "GodYao"
	maps["age"] = "28"
	fmt.Println(maps)

	// key exist
	if v, ok := maps["address"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("No Address")
	}

	// panic: assignment to entry in nil map
	nilMap := NilMapPanic()
	nilMap["data"] = "Go"

	fmt.Println(nilMap)
}
