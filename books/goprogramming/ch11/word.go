package ch11

func IsPalindrome(words string) bool {
	l := len(words)
	for i := range words {
		if words[i] != words[l-1-i] {
			return false
		}
	}
	return true
}
