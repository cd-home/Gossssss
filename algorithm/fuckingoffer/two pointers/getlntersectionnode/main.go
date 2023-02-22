package main

// 剑指 Offer 52. 两个链表的第一个公共节点

/*
输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
输出：Reference of the node with value = 8
输入解释：相交节点的值为 8 （注意，如果两个列表相交则不能为 0）。从各自的表头开始算起，
	链表 A 为 [4,1,8,4,5]
	链表 B 为 [5,0,1,8,4,5]。
	在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1 := headA
	p2 := headB
	for p1 != p2 {
		if p1 != nil {
			p1 = p1.Next
		} else {
			// p1 走完A, 回到B头继续
			p1 = headB
		}
		if p2 != nil {
			p2 = p2.Next
		} else {
			// p2 走完B, 回到A头继续
			p2 = headA
		}
	}
	return p1
}

func main() {

}
