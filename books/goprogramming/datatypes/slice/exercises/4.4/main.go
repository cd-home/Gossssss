package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(Rotate(s, 5))
	fmt.Println(s)
}

func Rotate[v any](s []v, p int) []v {
	r := s[p:]
	// 从 p 的位置往回遍历
	for i := p - 1; i >= 0; i-- {
		fmt.Printf("r=%v, len=%d, cap=%d\n", r, len(r), cap(r))
		r = append(r, s[i])
	}
	return r
}
