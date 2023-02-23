package main

import "strings"

// 剑指 Offer 58 - I. 翻转单词顺序

/*
输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。
为简单起见，标点符号和普通字母一样处理。
例如输入字符串"I am a student. "，则输出"student. a am I"。
示例 1：

	输入: "the sky is blue"
	输出: "blue is sky the"
*/
func reverseWords(s string) string {
	words := strings.Fields(s)
	i, j := 0, len(words)-1
	for i < j {
		words[i], words[j] = words[j], words[i]
		i++
		j--
	}
	return strings.Join(words, " ")
}
func main() {

}
