package main

// 剑指 Offer 50. 第一个只出现一次的字符

/*
	在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。
	示例 1:
		输入：s = "abaccdeff"
		输出：'b'
*/

func firstUniqChar(s string) byte {
	if s == "" {
		return ' '
	}
	m := make(map[rune]bool)
	// map 保存是否是出现一次
	for _, char := range s {
		// 从未出现, 写入true
		if _, ok := m[char]; !ok {
			m[char] = true
			// 再次出现修改为false
		} else {
			m[char] = false
		}
	}
	// 遍历是有顺序的
	for _, char := range s {
		if ok, _ := m[char]; ok {
			return byte(char)
		}
	}
	return ' '
}

func main() {

}
