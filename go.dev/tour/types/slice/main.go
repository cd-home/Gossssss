package main

import "fmt"

func printSlice(s []int) {
	// The length of a slice is the number of elements it contains.
	// The capacity of a slice is the number of elements in the underlying array,
	// counting from the first element in the slice
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func DeclareSlice() {
	fmt.Println("slice literal")
	// A slice literal is like an array literal without the length.
	// 字面量
	slices2 := []int{45, 46, 47}
	printSlice(slices2)

	fmt.Println("nil slice")
	// nil slice
	// A nil slice has a length and capacity of 0 and has no underlying array.
	var nilSlice []int
	// you can append
	nilSlice = append(nilSlice, 1)
	printSlice(nilSlice)

	// make 可以用来分配 带有长度 与 容量的 切片
	// 通常分配预估大小与容量的切片是友好的
	fmt.Println("make slice")
	// make
	slices := make([]int, 2, 10)
	slices = append(slices, 1)
	slices = append(slices, 2)
	slices = append(slices, 3)
	printSlice(slices)
}

func LenCapSlice() {
	fmt.Println("LenCapSlice")
	var slices []int
	printSlice(slices)

	slices = append(slices, 1)
	printSlice(slices)

	slices = append(slices, 2)
	printSlice(slices)

	// cap double
	slices = append(slices, 3)
	printSlice(slices)
}

func SubSlice() {
	fmt.Println("SubSlice")
	slices := []int{1, 2, 3, 4}
	// len = 2, cap = 3, count first to end
	printSlice(slices)
	sub1 := slices[1:3]
	printSlice(sub1)
	sub2 := slices[:]
	printSlice(sub2)

	// A slice does not store any data, it just describes a section of an underlying array.
	// share the same underlying array
	fmt.Println("share")
	sub1[0] = 9
	printSlice(slices)
	printSlice(sub1)
	printSlice(sub2)

	fmt.Println("over")
	// over => new slice
	// sub1 是子切片与 slices 共用底层数组, 但是当容量不够的时候, sub1扩容就会搬移到新的底层数组上面
	sub1 = append(sub1, []int{5, 6, 7, 8, 9}...)
	// 扩容之后再修改sub1, 就会发现已经是新的底层数组了
	sub1[0] = 10
	printSlice(slices)
	printSlice(sub1)
	printSlice(sub2)
}

func SubSlicing() {
	fmt.Println("SubSlicing")
	slices := []int{1, 2, 3, 4}

	printSlice(slices[:])
	printSlice(slices[:2])
	printSlice(slices[3:])
	printSlice(slices[1:3])
}

// ForRangeSlice
// for-range 可用来遍历容器类型 如 数组、切片、映射、通道
func ForRangeSlice() {
	fmt.Println("ForRangeSlice")
	slices := []int{1, 2, 3, 4}
	// 第一个是索引, 第二个是对应索引值的复制(特别注意)
	for index, value := range slices {
		fmt.Println(index, value)
		slices[index] *= 2
	}
	printSlice(slices)
}

func AppendSlices() {
	fmt.Println("AppendSlices")
	// A nil slice has a length and capacity of 0 and has no underlying array.
	// nil slice 没有长度与容量=0, 并且没有底层的数组
	// nil slice can append
	var slices []int
	slices = append(slices, 1)
	slices = append(slices, 2)
	others := []int{3, 4, 5}

	// can append slices
	// append 支持添加多个 (参数为不定参数模式)
	slices = append(slices, others...)
	slices = append(slices, 7, 8, 9)
	printSlice(slices)
}

func DeleteByIndex(slice []int, index int) []int {
	fmt.Println("DeleteByIndex")
	return append(slice[:index], slice[index+1:]...)
}

// slice可以通过定义, 可以通过获取数组或者切片的子切片得到
func main() {
	DeclareSlice()
	LenCapSlice()
	SubSlice()
	ForRangeSlice()
	AppendSlices()

	slices := []int{1, 2, 3, 4}
	printSlice(DeleteByIndex(slices, 2))

	SubSlicing()
}
