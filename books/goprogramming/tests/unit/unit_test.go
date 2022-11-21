package unit

import (
	"testing"
	"unicode"
)

func TestIsPalindromeBasic(t *testing.T) {
	// Basic
	if !IsPalindrome("abba") {
		t.Error("false")
	}
	if IsPalindrome("abb") {
		t.Error("false")
	}
}

func TestIsPalindromeTable(t *testing.T) {
	// 表测试
	tests := []struct {
		// 输入
		input string
		// 想要得到的结果
		want bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"A man, a plan, a canal: Panama", true},
	}
	for _, test := range tests {
		// got ?= want
		t.Run("", func(tt *testing.T) {
			if got := IsPalindrome(test.input); got != test.want {
				// 测试报告; 亦可以详细一点
				tt.Errorf("IsPalindrome(%q) = %v", test.input, got)
			}
		})
	}
}

func IsPalindrome(s string) bool {
	var letters []rune
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
