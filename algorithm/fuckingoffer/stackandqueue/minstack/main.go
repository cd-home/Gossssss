package minstack

// 剑指 Offer 30. 包含min函数的栈

/*
	难度： 简单
	定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中
	调用 min、push 及 pop 的时间复杂度都是 O(1)。
*/

type MinStack struct {
	stack []int
	min   []int
}

func ConstructorStack() MinStack {
	return MinStack{
		stack: make([]int, 0),
		min:   make([]int, 0),
	}
}

func (ms *MinStack) Push(x int) {
	// stack 正常 push
	ms.stack = append(ms.stack, x)
	// 如果 min stack 为空
	if len(ms.min) == 0 {
		ms.min = append(ms.min, x)
	} else {
		// stack 每一个元素层级(同期)都会对应一个最小值
		// 所以在pop的时候可以无差别pop
		min := ms.Min()
		if x < min {
			ms.min = append(ms.min, x)
		} else {
			ms.min = append(ms.min, min)
		}
	}
}

func (ms *MinStack) Pop() {
	// 由于 min 与 stack 都是对应同级, 全部正常pop
	ms.stack = ms.stack[:len(ms.stack)-1]
	ms.min = ms.min[:len(ms.min)-1]
}

func (ms *MinStack) Top() int {
	return ms.stack[len(ms.stack)-1]
}

func (ms *MinStack) Min() int {
	return ms.min[len(ms.min)-1]
}
