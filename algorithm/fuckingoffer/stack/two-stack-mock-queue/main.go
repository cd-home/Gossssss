package main

// 剑指 Offer 09. 用两个栈实现队列

/*
	难度: 简单
	用两个栈实现一个队列
	队列的声明如下:
		请实现它的两个函数 appendTail 和 deleteHead
		分别完成在队列尾部插入整数和在队列头部删除整数的功能
		(若队列中没有元素，deleteHead操作返回 -1 )
*/

type CQueue struct {
	InStack  []int
	OutStack []int
}

func Constructor() CQueue {
	return CQueue{
		InStack:  make([]int, 0),
		OutStack: make([]int, 0),
	}
}

func (cq CQueue) AppendTail(value int) {
	cq.InStack = append(cq.InStack, value)
}

func (cq CQueue) DeleteHead() int {
	// 直接特殊情况 -1
	if len(cq.InStack) == 0 && len(cq.OutStack) == 0 {
		return -1
	}
	// OutStack 有的情况下
	if len(cq.OutStack) > 0 {
		value := cq.OutStack[len(cq.OutStack)-1]
		// 要删除value
		cq.OutStack = cq.OutStack[:len(cq.OutStack)-1]
		return value
	}
	// OutStack 为空, 但是 InStack 有的情况下
	if len(cq.OutStack) == 0 && len(cq.InStack) > 0 {
		// 搬移 InStack => OutStack
		for len(cq.InStack) > 0 {
			item := cq.InStack[len(cq.InStack)-1]
			cq.InStack = cq.InStack[:len(cq.InStack)-1]
			cq.OutStack = append(cq.OutStack, item)
		}
	}
	// 搬移完成从 OutStack 返回
	value := cq.OutStack[len(cq.OutStack)-1]
	cq.OutStack = cq.OutStack[:len(cq.OutStack)-1]
	return value
}
