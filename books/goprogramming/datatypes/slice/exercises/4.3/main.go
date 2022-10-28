package main

import "fmt"

func main() {
	// 数组, 传递数组指针
	arr := [5]int{1, 2, 3, 4, 5}
	Reverse(&arr)
	fmt.Println(arr)

	arr2 := [5]string{"a", "b", "c", "d", "e"}
	Reverse(&arr2)
	fmt.Println(arr2)
}

// Reverse
// 数组指针, 无法动态的传递, 很麻烦
func Reverse[v any](arr *[5]v) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
