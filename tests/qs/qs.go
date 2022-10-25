package main

import "fmt"

func main() {
	arr := []int{2, 4, 10, 20, 1, 90}
	Qs(arr)
	fmt.Println(arr)
}

func Qs(arr []int) {
	qs(arr, 0, len(arr)-1)
}

func qs(arr []int, left, right int) {
	if left >= right {
		return
	}
	i := partition(arr, left, right)
	qs(arr, left, i-1)
	qs(arr, i+1, right)
}

func partition(arr []int, left, right int) int {
	pivot := arr[left]
	for left < right {
		for left < right && arr[right] > pivot {
			right--
		}
		arr[left] = arr[right]
		for left < right && arr[left] < pivot {
			left++
		}
		arr[right] = arr[left]
	}
	arr[left] = pivot
	return left
}
