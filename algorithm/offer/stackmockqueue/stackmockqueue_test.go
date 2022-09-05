package stackmockqueue

import (
	"testing"
)

// 剑指 Offer 09. 用两个栈实现队列
/*

	用两个栈实现一个队列。队列的声明如下，
	请实现它的两个函数 appendTail 和 deleteHead ，
	分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )
*/
type CQueue struct {
	In  []int
	Out []int
}

func Constructor() CQueue {
	return CQueue{
		In:  make([]int, 0),
		Out: make([]int, 0),
	}
}

func (c *CQueue) appendTail(value int) *CQueue {
	c.In = append(c.In, value)
	return c
}

func (c *CQueue) deleteHead() int {
	if len(c.In) == 0 && len(c.Out) == 0 {
		return -1
	}
	if len(c.Out) > 0 {
		value := c.Out[len(c.Out)-1]
		c.Out = c.Out[:len(c.Out)-1]
		return value
	}
	for len(c.In) > 0 {
		value := c.In[len(c.In)-1]
		c.In = c.In[:len(c.In)-1]
		c.Out = append(c.Out, value)
	}
	value := c.Out[len(c.Out)-1]
	c.Out = c.Out[:len(c.Out)-1]
	return value
}

func TestStackQueue(t *testing.T) {
	queue := Constructor()

	queue.appendTail(1).appendTail(2).appendTail(3)

	t.Log(queue.deleteHead())
	t.Log(queue.deleteHead())
	t.Log(queue.deleteHead())
	// empty
	t.Log(queue.deleteHead())
}
