package x

import "testing"

func TestNilSlice(t *testing.T) {
	var sliceTest []int
	t.Helper()
	if sliceTest == nil {
		t.Log("silceTest == nil")
	} else {
		t.Errorf("sliceTest != nil")
	}
	// nil slice can be append
	sliceTest = append(sliceTest, 1)
	t.Log(sliceTest)
}

func TestMakeSlice(t *testing.T) {
	// Make a slice
	sliceTest := make([]int, 10)
	t.Log(sliceTest)
	t.Helper()
	if sliceTest != nil {
		t.Log("silceTest != nil")
	} else {
		t.Errorf("sliceTest == nil")
	}
}

func TestSubslice(t *testing.T) {
	t.Helper()
	var sliceTest []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	t.Log(sliceTest)
	t.Log(len(sliceTest))
	t.Log(cap(sliceTest))

	// [ ) mode
	ss := sliceTest[1:3]
	t.Log(ss)

	// share array (pointer)
	ss[0] = 10
	t.Log(sliceTest)
	t.Log(ss)

	// cap len
	t.Log(len(sliceTest))
	t.Log(len(ss))

	t.Log(cap(sliceTest))
	t.Log(cap(ss))
}

func TestCapSlices(t *testing.T) {
	var sliceTest []int

	sliceTest = append(sliceTest, 1)
	t.Log(sliceTest)
	t.Log(len(sliceTest))
	t.Log(cap(sliceTest)) // 1

	sliceTest = append(sliceTest, 2)
	t.Log(sliceTest)
	t.Log(len(sliceTest))
	t.Log(cap(sliceTest)) // 2 double

	sliceTest = append(sliceTest, 3)
	t.Log(sliceTest)
	t.Log(len(sliceTest))
	t.Log(cap(sliceTest)) // 4 double
}

func TestCopySlices(t *testing.T) {
	var sliceTest []int
	sliceTest = append(sliceTest, 1, 2, 3)

	t.Log(sliceTest)

	// can't copy slices without len
	var copySlice []int
	copy(copySlice, sliceTest)

	t.Log(copySlice)

	// need give len(sliceTest) to copy
	_copySlice := make([]int, len(sliceTest))
	copy(_copySlice, sliceTest)
	t.Log(_copySlice)

	// _copySlice is new slice
	_copySlice[0] = 100
	t.Log(_copySlice)
	t.Log(sliceTest)
}

func TestSliceAppend(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	s1 = append(s1, s2...)
	t.Log(s1)
}
