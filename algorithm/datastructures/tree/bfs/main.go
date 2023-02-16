package main

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

// levelOrder 层序遍历 广度搜索
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	// 队列
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		var temp []int
		for i := 0; i < len(queue); i++ {
			node := queue[0]
			queue = queue[1:]
			temp = append(temp, node.Data)
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
