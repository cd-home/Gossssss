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

func main() {

}
