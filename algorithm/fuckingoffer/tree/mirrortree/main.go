package main

// 剑指 Offer 27. 二叉树的镜像

/*
	请完成一个函数，输入一个二叉树，该函数输出它的镜像。
	例如输入：

	  4
	 /  \
	 2   7
	/ \  / \
	1  3 6  9
	镜像输出：

	  4
	 /  \
	 7   2
	/ \  / \
	9  6 3 1
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 先找到叶子节点, 从下往上交换左右孩子
	left := mirrorTree(root.Left)
	right := mirrorTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}

func mirrorTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var stack = make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		node.Left, node.Right = node.Right, node.Left
	}
	return root
}

func main() {

}
