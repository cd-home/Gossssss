package main

// 剑指 Offer 28. 对称的二叉树

/*
请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。
例如，二叉树[1,2,2,3,4,4,3] 是对称的。
	 1
	/ \
   2   2
 /  \  / \
3  4  4   3

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	var check func(*TreeNode, *TreeNode) bool
	check = func(p, q *TreeNode) bool {
		// 都为nil情况
		if p == nil && q == nil {
			return true
		}
		// 其中一个为nil情况
		if p == nil || q == nil {
			return false
		}
		return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
	}
	return check(root, root)
}

func main() {

}
