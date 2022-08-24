package main

import "fmt"

// MakeMap
// The make function returns a map of the given type,
// initialized and ready for use.
func MakeMap(cap uint8) map[string]string {
	dataMap := make(map[string]string, cap)
	return dataMap
}

// NilMapPanic
// The zero value of a map is nil
func NilMapPanic() map[string]string {
	var maps map[string]string
	return maps
}

func ForRangeMap() {
	fmt.Println("ForRangeMap")
	maps := MakeMap(10)
	maps["name"] = "GodYao"
	maps["age"] = "28"
	for k, v := range maps {
		fmt.Println(k, v)
	}
}

func RetrieveElement() {
	fmt.Println("RetrieveElement")
	maps := MakeMap(10)
	maps["name"] = "GodYao"
	maps["age"] = "28"
	fmt.Println(maps)
	// retrieve value
	fmt.Println(maps["name"])

	// key exist
	// If key is not in the map, then elem is the zero value for the map's element type.
	if v, ok := maps["address"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("No key: Address")
	}
}

func DeleteElement() {
	fmt.Println("DeleteElement")
	maps := MakeMap(10)
	maps["name"] = "GodYao"
	maps["age"] = "28"
	fmt.Println(maps)

	delete(maps, "name")

	fmt.Println(maps)
}

// A map maps keys to values.
func main() {
	// make map
	m := MakeMap(10)
	fmt.Println(m)

	// panic: assignment to entry in nil map
	//nilMap := NilMapPanic()
	//nilMap["data"] = "Go"
	//
	//fmt.Println(nilMap)

	RetrieveElement()

	ForRangeMap()

	DeleteElement()

}
