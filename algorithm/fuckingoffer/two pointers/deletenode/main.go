package main

// 剑指 Offer 18. 删除链表的节点

/*
	给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
	返回删除后的链表的头节点。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}
	cur := head
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
			break
		}
		cur = cur.Next
	}
	return head
}

func main() {

}
