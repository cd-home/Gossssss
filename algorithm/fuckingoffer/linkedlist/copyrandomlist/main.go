package main

// 剑指 Offer 35. 复杂链表的复制

/*
	请实现 copyRandomList 函数，复制一个复杂链表。
	在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，
	还有一个 random 指针指向链表中的任意节点或者 null
*/

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	var m = make(map[*Node]*Node)
	var cur = head
	// copy 所有的节点, 正常Next的节点即是全部节点
	for cur != nil {
		m[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		// 原节点对应新建的节点
		node := m[cur]
		// 新节点的Next, 在map中对应 原节点cur的Next 的新节点
		node.Next = m[cur.Next]
		// 同Next
		node.Random = m[cur.Random]
	}
	// 原head节点在map中对应新的head节点
	return m[head]
}

func copyRandomList2(head *Node) *Node {
	if head == nil {
		return nil
	}
	// 构造新的链表 A->A1->B->B1->C->C1
	var cur = head
	for cur != nil {
		node := &Node{Val: cur.Val}
		node.Next = cur.Next
		cur.Next = node
		// cur 要指向原始节点
		cur = node.Next
	}
	// 处理 random 问题
	cur = head
	for cur != nil {
		if cur.Random != nil {
			// 新节点的 random 指向 当前节点随机节点的新节点 cur.Random.Next
			cur.Next.Random = cur.Random.Next
		}
		// 遍历的是原始节点, 要跳一个
		cur = cur.Next.Next
	}
	// 原链表
	var pre = head
	// 新链表
	var res = head.Next
	cur = head.Next
	for cur.Next != nil {
		// 拆分链表
		pre.Next = pre.Next.Next
		cur.Next = cur.Next.Next
		pre = pre.Next
		cur = cur.Next
	}
	// 处理原链表最后一个指向了新的节点问题
	pre.Next = nil
	return res
}

func main() {

}
