package main

import "fmt"

func main() {
	s := []string{"a", "b", "b", "c", "d", "d", "d", "d", "d", "e", "f"}
	ss := RemoveMulti(s)
	fmt.Println(ss)
}

// RemoveMulti
// 删除相同的相邻元素
func RemoveMulti[v comparable](s []v) []v {
	l := len(s)
	for i := 0; i < l-1; i++ {
		prev := s[i]
		next := s[i+1]
		if prev == next {
			s = append(s[:i], s[i+1:]...)
			// d d d 的情况就有问题了, 我回退一次, 避免后面的还有相同的
			i--
			l--
		}
	}
	return s[:len(s)-1]
}
