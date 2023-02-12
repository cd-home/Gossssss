package main

// 剑指 Offer 24. 反转链表

/*
	定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	var cur = head
	for cur != nil {
		nexts := cur.Next
		cur.Next = pre
		pre = cur
		cur = nexts
	}
	return pre
}

func main() {

}
