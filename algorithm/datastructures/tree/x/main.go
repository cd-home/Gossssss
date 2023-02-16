package main

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

func preOrderTraversal(root *TreeNode) []int {
	var res []int
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	// root 判断可以拿出去
	for root != nil && len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Data)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return res
}

func inOrderTraversal(root *TreeNode) []int {
	var res []int
	stack := make([]*TreeNode, 0)
	cur := root
	for cur != nil || len(stack) > 0 {
		// 不停的找到最左
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Data)
		if node.Right != nil {
			cur = node.Right
		}
	}
	return res
}

func postOrderTraversal(root *TreeNode) []int {
	var temp []int
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	//  中 右 左
	for root != nil && len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		temp = append(temp, node.Data)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	res := make([]int, 0, len(temp))
	for len(temp) > 0 {
		res = append(res, temp[len(temp)-1])
		temp = temp[:len(temp)-1]
	}
	return res
}

func main() {

}
