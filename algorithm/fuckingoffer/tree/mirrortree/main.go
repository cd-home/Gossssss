package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 从下往上, 先找到叶子节点
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
