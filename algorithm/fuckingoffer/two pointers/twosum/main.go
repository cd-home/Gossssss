package main

// 剑指 Offer 57. 和为s的两个数字

/*
	输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。如果有多对数字的和等于s，则输出任意一对即可。
*/

func twoSum(nums []int, target int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		// 避免使用 nums[i] + nums[j] 判断, 可能溢出
		if target-nums[i] > nums[j] {
			i++
		} else if target-nums[i] < nums[j] {
			j--
		} else if target-nums[i] == nums[j] {
			return []int{nums[i], nums[j]}
		}
	}
	return nil
}

func main() {

}
