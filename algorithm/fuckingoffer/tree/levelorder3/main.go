package main

// 剑指 Offer 32 - III. 从上到下打印二叉树 III

/*
	请实现一个函数按照之字形顺序打印二叉树，
	即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，
	其他行以此类推。
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
	  [20,9],
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
		// 如果是奇数层就倒序 root
		if len(res)%2 == 1 {
			reverse(temp)
		}
		res = append(res, temp)
	}
	return res
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {

}
