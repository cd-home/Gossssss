package main

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

// preOrderTraversal 先序遍历
func preOrderTraversal(root *TreeNode) []int {
	// 返回的结果
	var res []int
	// 声明递归函数
	var preOrder func(*TreeNode, *[]int)
	preOrder = func(node *TreeNode, res *[]int) {
		if node != nil {
			*res = append(*res, node.Data)
			preOrder(node.Left, res)
			preOrder(node.Right, res)
		}
	}
	preOrder(root, &res)
	return res
}

// inOrderTraversal 中序遍历
func inOrderTraversal(root *TreeNode) []int {
	var res []int
	var inOrder func(*TreeNode, *[]int)
	inOrder = func(node *TreeNode, res *[]int) {
		inOrder(node.Left, res)
		*res = append(*res, node.Data)
		inOrder(node.Right, res)
	}
	inOrder(root, &res)
	return res
}

// postOrderTraversal 后序遍历
func postOrderTraversal(root *TreeNode) []int {
	var res []int
	var postOrder func(*TreeNode, *[]int)
	postOrder = func(node *TreeNode, res *[]int) {
		postOrder(node.Left, res)
		postOrder(node.Right, res)
		*res = append(*res, node.Data)
	}
	postOrder(root, &res)
	return res
}

func main() {

}
