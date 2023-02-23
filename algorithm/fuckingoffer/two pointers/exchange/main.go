package main

// 剑指 Offer 21. 调整数组顺序使奇数位于偶数前面

/*
	输入一个整数数组，实现一个函数来调整该数组中数字的顺序，
	使得所有奇数在数组的前半部分，所有偶数在数组的后半部分。
	% 2 可以判断奇数、偶数 ==>  &1 也可以判断
*/

func exchange(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]&1 == 0 && nums[j]&1 == 1 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
		// 这里可以用for方便判断后续都是奇数情况, 也可以不用
		for i < j && nums[i]&1 == 1 {
			i++
		}
		for i < j && nums[j]&1 == 0 {
			j--
		}
	}
	return nums
}

func main() {

}
