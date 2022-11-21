package main

import (
	"fmt"
	"unicode"
)

func main() {
	ExampleIsPalindrome()
}

func ExampleIsPalindrome() {
	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(IsPalindrome("palindrome"))
	// Output:
	// true
	// false
}

func IsPalindrome(s string) bool {
	// 预先声明避免多次增长分配
	letters := make([]rune, 0, len(s))
	for _, r := range s {
		// 处理非字符情况
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	// 判断是否是回文
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
