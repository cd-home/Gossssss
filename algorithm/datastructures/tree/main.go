package main

type Tree struct {
	Data  int
	Left  *Tree
	Right *Tree
}

// preOrderTraversal 先序遍历
func preOrderTraversal(root *Tree) []int {
	// 返回的结果
	var res []int
	// 声明递归函数
	var preOrder func(*Tree, *[]int)
	preOrder = func(node *Tree, res *[]int) {
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
func inOrderTraversal(root *Tree) []int {
	var res []int
	var inOrder func(*Tree, *[]int)
	inOrder = func(node *Tree, res *[]int) {
		inOrder(node.Left, res)
		*res = append(*res, node.Data)
		inOrder(node.Right, res)
	}
	inOrder(root, &res)
	return res
}

// postOrderTraversal 后序遍历
func postOrderTraversal(root *Tree) []int {
	var res []int
	var postOrder func(*Tree, *[]int)
	postOrder = func(node *Tree, res *[]int) {
		postOrder(node.Left, res)
		postOrder(node.Right, res)
		*res = append(*res, node.Data)
	}
	postOrder(root, &res)
	return res
}

func main() {

}
