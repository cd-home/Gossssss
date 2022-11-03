package main

import "fmt"

// Tree
// 二叉查找树(二叉搜索树)  left <= root < right
type Tree struct {
	root *node
}

// tree node
type node struct {
	val   int
	left  *node
	right *node
}

func (t *Tree) Add(val int) {
	if t.root == nil {
		t.root = &node{val: val}
		return
	}
	// 每次添加都从root开始判断,二分
	t.root.add(val)
}

func (n *node) add(val int) {
	// 左子树
	if val <= n.val {
		// 如果左子树存在, 递归
		if n.left != nil {
			n.left.add(val)
		} else {
			n.left = &node{val: val}
		}
		return
	}
	// 右子树
	if val > n.val {
		// 如果右子树存在, 递归
		if n.right != nil {
			n.right.add(val)
		} else {
			n.right = &node{val: val}
		}
		return
	}
}

func (t *Tree) Find(val int) *node {
	if t.root == nil {
		return nil
	}
	return t.root.find(val)
}

func (n *node) find(val int) *node {
	if n == nil {
		return nil
	}
	if val == n.val {
		return n
	}
	// 二分查找
	if val < n.val {
		return n.left.find(val)
	}
	if val > n.val {
		return n.right.find(val)
	}
	return nil
}

func (t *Tree) FindMin() *node {
	if t.root == nil {
		return nil
	}
	return t.root.findMin()
}

func (n *node) findMin() *node {
	if n == nil {
		return nil
	}
	min := n
	for min.left != nil {
		min = min.left
	}
	return min
}

func (t *Tree) FindMax() *node {
	if t.root == nil {
		return nil
	}
	return t.root.findMax()
}

func (n *node) findMax() *node {
	if n == nil {
		return nil
	}
	max := n
	for max.right != nil {
		max = max.right
	}
	return max
}

func (t *Tree) inorderTraversal() []int {
	var res []int
	var inOrder func(*node, *[]int)
	inOrder = func(t *node, res *[]int) {
		if t != nil {
			inOrder(t.left, res)
			*res = append(*res, t.val)
			inOrder(t.right, res)
		}
	}
	inOrder(t.root, &res)
	return res
}

func main() {
	arr := []int{9, 2, 2, 3, 10, 28, 1, 4}
	t := new(Tree)
	for _, v := range arr {
		t.Add(v)
	}
	fmt.Println(t.inorderTraversal())

	fmt.Println(t.Find(9))
	fmt.Println(t.FindMin())
	fmt.Println(t.FindMax())
}
