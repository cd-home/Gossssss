package main

import "fmt"

func main() {
	fmt.Println(fbDp(6))
	fmt.Println(fbR(6))
	fmt.Println(fbUltimate(6))
}

func fbDp(n int) int {
	if n < 2 {
		return n
	}
	f := make([]int, n+1)
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

func fbUltimate(n int) int {
	pp, p := 0, 1
	var cur int
	if n == 0 {
		return pp
	}
	if n == 1 {
		return p
	}
	for i := 2; i <= n; i++ {
		cur = p + pp
		pp = p
		p = cur
	}
	return cur
}

func fbR(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fbR(n-1) + fbR(n-2)
}
