package algssss

import (
	"testing"
)

func BubbleSort(arr []int) {
	l := len(arr)
	if l <= 1 {
		return
	}
	// 轮数
	for i := 0; i < l; i++ {
		// 每一轮做什么, 一个数和剩下的比较
		// 每一轮进来定义一个是否有数据交换的标志
		flag := false
		for j := 0; j < l-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

func MergeSort(arr []int) {
	l := len(arr)
	if l <= 1 {
		return
	}
	RecurMerge(arr, 0, l-1)
}

func RecurMerge(arr []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	RecurMerge(arr, start, mid)
	RecurMerge(arr, mid+1, end)
	Merge(arr, start, mid, end)
}

func Merge(arr []int, start, mid, end int) {
	tempArr := make([]int, end-start+1)
	i := start
	j := mid + 1
	k := 0
	for ; i <= mid && j <= end; k++ {
		if arr[i] <= arr[j] {
			tempArr[k] = arr[i]
			i++
		} else {
			tempArr[k] = arr[j]
			j++
		}
	}
	for ; i <= mid; i++ {
		tempArr[k] = arr[i]
		k++
	}
	for ; j <= end; j++ {
		tempArr[k] = arr[j]
		k++
	}
	copy(arr[start:end+1], tempArr)
}

func QuickSort(arr []int) {
	separateSort(arr, 0, len(arr)-1)
}

func separateSort(arr []int, start, end int) {
	if start >= end {
		return
	}
	i := partition(arr, start, end)
	separateSort(arr, start, i-1)
	separateSort(arr, i+1, end)
}

func partition(arr []int, start, end int) int {
	pivot := arr[end]
	var i = start
	// i移动+1，直到遇到比pivot大的数字,i标记了比pivot的数字
	// j一直往后遍历，找到比pivot小的 交换 i j
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]
	return i
}

func TestBubbleSort(t *testing.T) {

}
