package main

// 剑指 Offer 11. 旋转数组的最小数字

/*
	把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
	给你一个可能存在重复元素值的数组numbers，它原来是一个升序排列的数组，并按上述情形进行了一次旋转。
	请返回旋转数组的最小元素。例如，数组[3,4,5,1,2] 为 [1,2,3,4,5] 的一次旋转，
	该数组的最小值为 1。

*/

func minArray(numbers []int) int {
	left := 0
	right := len(numbers) - 1
	for left <= right {
		mid := (left + right) >> 1
		// mid 在旋转点的左侧数组 x [mid+1, j]
		if numbers[mid] > numbers[right] {
			left = mid + 1
			// mid 在旋转点的右侧数组 x [i, mid]
		} else if numbers[mid] < numbers[right] {
			right = mid
		} else {
			right--
		}
	}
	return numbers[left]
}

func main() {

}
