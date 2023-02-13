package main

// 剑指 Offer 05. 替换空格

/*
	请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
*/

func replaceSpace(s string) string {
	var res = ""
	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {
			res += "%20"
		} else {
			res += string(s[i])
		}
	}
	return res
}

func main() {

}
