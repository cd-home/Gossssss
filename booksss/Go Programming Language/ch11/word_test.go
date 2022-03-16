package ch11

import "testing"

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("sfdsgds") {
		t.Error("IsPalindrome('abba') = false")
	}
	if !IsPalindrome("abba") {
		t.Error("IsPalindrome('abba') = false")
	}
}
