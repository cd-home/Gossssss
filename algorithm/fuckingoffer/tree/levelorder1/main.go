package main

// 剑指 Offer 32 - I. 从上到下打印二叉树

/*
从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。
例如:

	给定二叉树:[3,9,20,null,null,15,7],
		3
	   / \
	  9  20
		/  \
	   15   7
	返回：[3,9,20,15,7]
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// levelOrder 层序遍历 广度搜索
func levelOrder(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		res = append(res, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return res
}

func main() {

}
