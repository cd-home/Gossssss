package main

import "fmt"

func foo() (ret int) {
	defer func() {
		ret++
	}()
	// 0 = ret, 然后在返回之前执行 defer ret++ ==> 1
	return 0
}

func foo2() int { //  0
	var i int
	defer func() {
		i++
	}()
	return i
}

func DeferReturn() (t int) { //1: 初始值为0  3: 被修改为2
	defer func(i int) {
		fmt.Println(i)
		fmt.Println(t) // t = 2
	}(t) // 2: 出现defer的时候参数就已经确定 t = 0 然后 i = 0
	t = 1
	return 2 // t = 2
}
func main() {
	//fmt.Println(foo())
	fmt.Println(foo2())
	//fmt.Println(DeferReturn())

	fmt.Println(DeferF2(1))
}

func DeferF1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return
}

func DeferF2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}
