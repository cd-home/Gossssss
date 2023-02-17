package main

// 剑指 Offer 26. 树的子结构

/*
	难度: 中等
	输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
	B是A的子结构， 即 A中有出现和B相同的结构和节点值。
	例如:
	给定的树 A:

			   3
			 / \
			4  5
		   / \
		  1  2
	给定的树 B：
		 4
		/
	   1
	返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	return hasSubStructure(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func hasSubStructure(A *TreeNode, B *TreeNode) bool {
	// B 为nil, 就已经递归到最后了, 说明是相同root的子结构
	if B == nil {
		return true
	}
	// 递归到A结束, 上面判断的B还不是nil; 或者节点值压根不同;
	if A == nil || A.Val != B.Val {
		return false
	}
	// 进入递归
	return hasSubStructure(A.Left, B.Left) && hasSubStructure(A.Right, B.Right)
}

func main() {

}
