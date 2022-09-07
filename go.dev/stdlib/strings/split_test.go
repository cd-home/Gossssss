package strings_test

import (
	"fmt"
	"strings"
	"testing"
)

func ReverseString(s string) string {
	sps := strings.Split(s, "")
	fmt.Println(sps)
	res := ""
	for i := len(sps) - 1; i >= 0; i-- {
		res += sps[i]
	}
	return res
}

func TestString(t *testing.T) {
	s := "我你他"
	fmt.Println(ReverseString(s))
}
