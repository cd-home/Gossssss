package main

// 剑指 Offer 32 - II. 从上到下打印二叉树 II

/*
	从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。
	例如:
	给定二叉树:[3,9,20,null,null,15,7],

		3
	   / \
	  9  20
		/  \
	   15   7
	返回其层次遍历结果：
	[
	  [3],
	  [9,20],
	  [15,7]
	]
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	var queue []*TreeNode
	if root == nil {
		return res
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		var temp []int
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			// 临时变量保存每一层的遍历结果
			temp = append(temp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, temp)
	}
	return res
}

func main() {

}
