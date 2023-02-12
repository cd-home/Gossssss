package reverselist

import "testing"

/**
 * Definition for singly-linked list
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var pre *ListNode
	cur := head
	for cur != nil {
		// keep nexts
		nexts := cur.Next

		// reverse
		cur.Next = pre

		// move pointer
		pre = cur
		cur = nexts
	}
	return pre
}

func TestReverseLinkList(t *testing.T) {
	ReverseList(&ListNode{})
}
