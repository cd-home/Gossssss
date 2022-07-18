package ch11

import "testing"

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("sfdsgds") {
		t.Log("IsPalindrome('abba') = false")
	}
	if !IsPalindrome("abba") {
		t.Log("IsPalindrome('abba') = false")
	}
}
