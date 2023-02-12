package main

// 剑指 Offer 06. 从尾到头打印链表

/*
	难度: 简单
	输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	var tmp []int
	var cur = head
	for cur != nil {
		tmp = append(tmp, cur.Val)
		cur = cur.Next
	}
	var res = make([]int, len(tmp))
	for i := len(tmp) - 1; i >= 0; i-- {
		res[len(tmp)-1-i] = tmp[i]
	}
	return res
}

func main() {

}
