package twosum

// 给定一个整数数组 nums 和一个整数目标值 target，
//请你在该数组中找出和为目标值 target 的那 两个整数，并返回它们的数组下标。

// twoSum
// 利用map存储value-index, 这样寻找更快(空间换时间)
func twoSum(nums []int, target int) []int {
	store := make(map[int]int)
	for idx, val := range nums {
		diff := target - val
		if i, ok := store[diff]; ok {
			return []int{idx, i}
		}
		store[val] = idx
	}
	return []int{-1, -1}
}
