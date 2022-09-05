package main

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

func main() {

}
