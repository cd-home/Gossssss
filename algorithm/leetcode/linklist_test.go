package algssss

import (
	"testing"
)

// 节点
type Node struct {
	Data interface{}
	Next *Node
}

// 链表
type LinkList struct {
	head   *Node //头节点，方便操作链表
	length uint
}

// 新建节点
func NewNode(data interface{}) *Node {
	return &Node{Data: data, Next: nil}
}

// 获取节点的下一个节点
func (n *Node) GetNodeNext() *Node {
	return n.Next
}

// 获取节点的值
func (n *Node) GetNodeData() interface{} {
	return n.Data
}

// 新建链表
func NewLinkList(data interface{}) *LinkList {
	return &LinkList{head: NewNode(data), length: 0}
}

// IsEmpty
func (ll *LinkList) IsEmpty() bool {
	return ll.head.Next == nil
}

// 反转链表
func (ll *LinkList) ReverseLinkList() bool {
	if ll.IsEmpty() {
		return false
	}
	cur := ll.head.Next
	var pre *Node
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	ll.head.Next = pre
	return true
}

// 反转链表2
func (ll *LinkList) ReverseLinkList2() bool {
	if ll.head == nil || ll.head.Next == nil {
		return false
	}
	// 迭代反转链表
	var pre *Node
	cur := ll.head.Next
	for cur != nil {
		nexts := cur.Next

		// pre 指向 cur
		cur.Next = pre
		
		// move
		pre = cur
		cur = nexts
	}
	ll.head.Next = pre
	return true
}

// 判断是否有环
func (ll *LinkList) HasCycle() bool {
	if ll.head != nil {
		// 快慢指针
		slow := ll.head
		fast := ll.head
		// 如果有环一定相遇，如果没有环 fast会跑到nil
		for fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
			if slow == fast {
				return true
			}
		}
	}
	return false
}

// 获取中间节点
func (ll *LinkList) FindMiddleNode() *Node {
	if ll.head == nil || ll.head.Next == nil {
		return nil
	}
	if ll.head.Next.Next == nil {
		return ll.head.Next
	}
	// 当fast走到最后的时候 slow刚好走到一半
	slow, fast := ll.head, ll.head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 获取倒数第K个节点
func (ll *LinkList) FindDescKNode(k uint) *Node {
	if k <= 0 || ll.head == nil || ll.length < k {
		return nil
	}
	first, second := ll.head.Next, ll.head.Next
	// 倒数第k个数，相距k - 1个位置
	k = k - 1
	for k > 0 {
		first = first.Next
		if first == nil {
			return nil
		}
		k--
	}
	for first.Next != nil {
		second = second.Next
		first = first.Next
	}
	return second
}

// 合并有序链表
func (ll *LinkList) MergeSortedLinkList(l *LinkList) *LinkList {
	if l == nil || l.head.Next == nil {
		return ll
	}
	if ll == nil || ll.head.Next == nil {
		return l
	}
	newLinkList := NewLinkList("Head")
	cur := newLinkList.head
	cur1 := ll.head.Next
	cur2 := l.head.Next
	for cur1 != nil && cur2 != nil {
		if cur1.Data.(int) > cur2.Data.(int) {
			cur.Next = cur2
			cur2 = cur2.Next
		} else {
			cur.Next = cur1
			cur1 = cur1.Next
		}
		cur = cur.Next
	}
	// 有一个curX为nil，中断循环，后面就可以直接添加到cur后面
	if cur1 != nil {
		cur.Next = cur1
	} else if cur2 != nil {
		cur.Next = cur2
	}
	return newLinkList
}

func TestReverseLinkList(t *testing.T) {

}
