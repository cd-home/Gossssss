package main

import "fmt"

func DeclareSlice() {
	var nilSlice []int
	// you can append
	nilSlice = append(nilSlice, 1)
	fmt.Println(nilSlice)

	// make
	slices := make([]int, 2, 10)
	slices = append(slices, 1)
	slices = append(slices, 2)
	slices = append(slices, 3)
	fmt.Println(slices)
}

func LenCapSlice() {
	fmt.Println("LenCapSlice")
	var slices []int
	fmt.Println(len(slices), cap(slices))

	slices = append(slices, 1)
	fmt.Println(len(slices), cap(slices))

	slices = append(slices, 2)
	fmt.Println(len(slices), cap(slices))

	// cap double
	slices = append(slices, 3)
	fmt.Println(len(slices), cap(slices))
}

func SubSlice() {
	fmt.Println("SubSlice")
	slices := []int{1, 2, 3, 4}
	sub1 := slices[1:3]
	fmt.Println(sub1)
	fmt.Println(len(sub1), cap(sub1))
	sub2 := slices[:]
	fmt.Println(sub2)

	// shadow
	sub1[0] = 9
	fmt.Println(slices)
	fmt.Println(sub1)
	fmt.Println(sub2)

	// over => new slice
	sub1 = append(sub1, []int{5, 6, 7, 8, 9}...)
	sub1[0] = 10
	fmt.Println(slices)
	fmt.Println(sub1)
	fmt.Println(sub2)
}

func ForRangeSlice() {
	fmt.Println("ForRangeSlice")
	slices := []int{1, 2, 3, 4}
	for index, value := range slices {
		fmt.Println(index, value)
		slices[index] *= 2
	}
	fmt.Println(slices)
}

func AppendSlices() {
	fmt.Println("AppendSlices")
	// nil slice can append
	var slices []int
	slices = append(slices, 1)
	slices = append(slices, 2)
	others := []int{3, 4, 5}
	// can append slices
	slices = append(slices, others...)

	fmt.Println(slices)
}

func DeleteByIndex(slice []int, index int) []int {
	fmt.Println("DeleteByIndex")
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	DeclareSlice()
	LenCapSlice()
	SubSlice()
	ForRangeSlice()
	AppendSlices()

	slices := []int{1, 2, 3, 4}
	fmt.Println(DeleteByIndex(slices, 2))
}
