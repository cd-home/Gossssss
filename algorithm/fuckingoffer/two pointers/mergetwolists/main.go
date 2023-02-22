package main

// 剑指 Offer 25. 合并两个排序的链表

/*
	输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

	示例1：
		输入：1->2->4, 1->3->4
		输出：1->1->2->3->4->4
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	} else if l2 != nil {
		cur.Next = l2
	}
	return dummy.Next
}
func main() {

}
