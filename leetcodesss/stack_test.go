package leetcodesss

type Stacker interface {
	Push(item interface{})
	Pop() interface{}
	IsEmpty() bool
	Top() interface{}
	Flush() bool
}

type Stack struct {
	data []interface{}
	top int
}

func NewStack(size int) *Stack {
	return &Stack{
		data: make([]interface{}, 0, size),
		top: -1,
	}
}

func (stack *Stack) IsEmpty() bool {
	return stack.top < 0
}

func (stack *Stack) Push(item interface{}) {
	stack.top += 1
	if stack.top > len(stack.data) - 1 {
		stack.data = append(stack.data, item)
	} else {
		stack.data[stack.top] = item
	}
}

func (stack *Stack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	item := stack.data[stack.top]
	stack.top--
	return item
}

func (stack *Stack) Top() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	item := stack.data[stack.top]
	stack.top--
	return item
}

func (stack *Stack) Flush() {
	stack.top = -1
}