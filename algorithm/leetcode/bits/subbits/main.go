package main

// 数组所有子集情况，不包括空集合 n^2 - 1
/*
	0. 假设有三个元素{1, 2, 3}
	1. n个元素，采用二进制表示是否在子集中如
		{}             000			0
		{3}            001          1
		{2} 		   010			2
		{2, 3}         011			3
		{1}            100			4
		{1, 3}         101			5
		{1, 2}         110			6
		{1, 2, 3}      111			7
	2. 01位置序列就是数组子集的序列，恰好01序列又对应2^n - 1数字
	3. 可以最终转换为求2^n - 1中所有的数二进制表示
*/

func SubSets(nums []uint) (ans [][]uint) {

	var n = uint(len(nums))
	var mask uint

	// 1 << n 表示 2^n
	for mask = 0; mask < (1 << n); mask++ {

		var set []uint
		var i int
		var v uint

		for i, v = range nums {
			// 判断当前的id二进制序列是否已经存在？
			if mask>>i&1 > 0 {
				set = append(set, v)
			}
		}
		// 返回所有的子集
		ans = append(ans, append([]uint(nil), set...))
	}

	return
}

func main() {

}
