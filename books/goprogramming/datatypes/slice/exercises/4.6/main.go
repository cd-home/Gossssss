package main

import (
	"fmt"
	"regexp"
	"unicode"
)

func main() {
	s := "aaaa   bvbb s ss   asa   sdsa"
	ss := removeEmpty(s)
	fmt.Println(ss)
	fmt.Println(s)

	fmt.Println(replaceEmpty(s))
}

func removeEmpty(str string) string {
	bs := []byte(str)
	l := len(bs)
	for i := 0; i < l-1; i++ {
		p := bs[i]
		q := bs[i+1]
		// 判断是否是空格
		if unicode.IsSpace(rune(p)) && unicode.IsSpace(rune(q)) {
			bs = append(bs[:i], bs[i+1:]...)
			i--
			l--
		}
	}
	return string(bs)
}

// 正则的方式匹配空格替换
func replaceEmpty(str string) string {
	//reg, err := regexp.CompilePOSIX(`[[:space:]]+`)
	reg := regexp.MustCompile(`\s+`)
	return reg.ReplaceAllString(str, " ")
}
