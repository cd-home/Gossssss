package main

// 剑指 Offer 04. 二维数组中的查找

/*
	在一个 n * m 的二维数组中，每一行都按照从左到右非递减的顺序排序，
	每一列都按照从上到下非递减的顺序排序。
	请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
示例:
	现有矩阵 matrix 如下：

	[
	  [1,   4,  7, 11, 15],
	  [2,   5,  8, 12, 19],
	  [3,   6,  9, 16, 22],
	  [10, 13, 14, 17, 24],
	  [18, 21, 23, 26, 30]
	]
*/

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	row, col := 0, len(matrix[0])-1
	for row <= len(matrix)-1 && col >= 0 {
		if matrix[row][col] == target {
			return true
		}
		if matrix[row][col] > target {
			col--
		} else {
			row++
		}
	}
	return false
}

func main() {

}
