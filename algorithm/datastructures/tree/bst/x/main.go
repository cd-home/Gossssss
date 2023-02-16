package main

import "fmt"

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

// FindX 循环
func FindX(node *TreeNode, data int) *TreeNode {
	for node != nil {
		if data < node.Data {
			node = node.Left
		} else if data > node.Data {
			node = node.Right
		} else {
			return node
		}
	}
	return nil
}

// FindY 递归
func FindY(node *TreeNode, data int) *TreeNode {
	if node == nil {
		return nil
	}
	if node.Data == data {
		return node
	}
	if data < node.Data {
		return FindY(node.Left, data)
	} else {
		return FindY(node.Right, data)
	}
}

func FindMaxX(node *TreeNode) *TreeNode {
	if node != nil {
		for node.Right != nil {
			node = node.Right
		}
	}
	return node
}

func FindMaxY(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	if node.Right != nil {
		return FindMaxY(node.Right)
	} else {
		return node
	}
}

func FindMinX(node *TreeNode) *TreeNode {
	if node != nil {
		for node.Left != nil {
			node = node.Left
		}
	}
	return node
}

func FindMinY(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	if node.Left != nil {
		return FindMinY(node.Left)
	} else {
		return node
	}
}

func Insert(node *TreeNode, data int) *TreeNode {
	if node == nil {
		node = &TreeNode{Data: data}
	} else {
		if data < node.Data {
			node.Left = Insert(node.Left, data)
		} else {
			node.Right = Insert(node.Right, data)
		}
	}
	return node
}

func Delete(node *TreeNode, data int) {

}

func InOrderTraversal(root *TreeNode) []int {
	var res []int
	var inOrder func(*TreeNode, *[]int)
	inOrder = func(node *TreeNode, res *[]int) {
		if node != nil {
			inOrder(node.Left, res)
			*res = append(*res, node.Data)
			inOrder(node.Right, res)
		}
	}
	inOrder(root, &res)
	return res
}

func main() {
	root := &TreeNode{Data: 8}
	Insert(root, 4)
	Insert(root, 5)
	Insert(root, 3)
	Insert(root, 9)
	Insert(root, 1)
	Insert(root, 2)
	Insert(root, 10)
	Insert(root, 17)

	// 中序遍历
	fmt.Println(InOrderTraversal(root))

	fmt.Println(FindMaxX(root))
	fmt.Println(FindMaxY(root))

	fmt.Println(FindMinX(root))
	fmt.Println(FindMinY(root))

	fmt.Println(FindX(root, 10))
	fmt.Println(FindY(root, 10))
}
