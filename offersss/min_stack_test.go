package offersss

import (
	"testing"
)

// 剑指 Offer 30. 包含min函数的栈

/*

	定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中
	调用 min、push 及 pop 的时间复杂度都是 O(1)。
*/

type MinStack struct {
	stack []int
	min   []int
}

/** initialize your data structure here. */
func ConstructorStack() MinStack {
	return MinStack{
		stack: make([]int, 0),
		min:   make([]int, 0),
	}
}

func (ms *MinStack) Push(x int) {
	ms.stack = append(ms.stack, x)
	if len(ms.stack) > 0 {
		m := ms.Min()
		// 注意需要同步增长，保证每一个添加的x, 在min中都有一个同时期的最小值
		if x < m {
			ms.min = append(ms.min, x)
		} else {
			ms.min = append(ms.min, m)
		}
	} else {
		ms.min = append(ms.min, x)
	}
}

func (ms *MinStack) Pop() {
	if len(ms.stack) > 0 {
		ms.stack = ms.stack[:len(ms.stack)-1]
		ms.min = ms.min[:len(ms.min)-1]
	}
}

func (ms *MinStack) Top() int {
	return ms.stack[len(ms.stack)-1]
}

func (ms *MinStack) Min() int {
	return ms.min[len(ms.min)-1]
}

func TestMinStack(t *testing.T) {
	stack := ConstructorStack()
	stack.Push(1)
	stack.Push(-1)
	stack.Push(-2)

	t.Log(stack.stack)
	t.Log(stack.min)

	t.Log(stack.Top())

	stack.Pop()

	t.Log(stack.stack)
	t.Log(stack.min)
}
