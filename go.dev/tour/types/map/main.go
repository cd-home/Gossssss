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
// The zero value of a map is nil, nil map 无法使用, 必须初始化
func NilMapPanic() map[string]string {
	var maps map[string]string
	return maps
}

func ForRangeMap() {
	fmt.Println("ForRangeMap")
	maps := MakeMap(10)
	maps["name"] = "yao"
	maps["age"] = "28"
	for k, v := range maps {
		fmt.Println(k, v)
	}
}

func RetrieveElement() {
	fmt.Println("RetrieveElement")
	maps := MakeMap(10)
	maps["name"] = "yao"
	maps["age"] = "28"

	fmt.Println(maps)

	// retrieve value
	name := maps["name"]
	fmt.Println(name)

	// key exist, ok 模式
	// If key is not in the map, then elem is the zero value for the map's element type.
	if v, ok := maps["address"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("No key: Address")
	}

	// 修改值, 无则添加, 有则修改
	maps["name"] = "mike"

	fmt.Println(maps)
}

func DeleteElement() {
	fmt.Println("DeleteElement")
	maps := MakeMap(10)
	maps["name"] = "yao"
	maps["age"] = "28"
	fmt.Println(maps)

	// 删除不存在的key 不会报错
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

	// map 字面量模式
	var user = map[string]string{
		"name": "yao",
	}
	fmt.Println(user)
}
