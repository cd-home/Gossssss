package bench

import (
	"testing"
	"unicode"
)

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
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
