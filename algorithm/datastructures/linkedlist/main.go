package main

// Node
// 节点
type Node struct {
	Data interface{}
	Next *Node
}

// LinkList
// 链表
type LinkList struct {
	head   *Node //头节点，方便操作链表
	length uint
}

// NewNode
// 新建节点
func NewNode(data interface{}) *Node {
	return &Node{Data: data, Next: nil}
}

// GetNodeNext
// 获取节点的下一个节点
func (n *Node) GetNodeNext() *Node {
	return n.Next
}

// GetNodeData
// 获取节点的值
func (n *Node) GetNodeData() interface{} {
	return n.Data
}

// NewLinkList
// 新建链表
func NewLinkList(data interface{}) *LinkList {
	return &LinkList{head: NewNode(data), length: 0}
}

// IsEmpty
// 是否为空
func (ll *LinkList) IsEmpty() bool {
	return ll.head.Next == nil
}

func main() {

}