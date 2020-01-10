package solutiion

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Runtime 0 ms, faster than 100%
// Memory 4.5 MB, less than 100%

var max int

func diameterOfBinaryTree(root *TreeNode) int {
	max = 1
	dfs(root)
	return max - 1
}

func dfs(node *TreeNode) int {
	if node == nil {
		return 0
	}
	l := dfs(node.Left)
	r := dfs(node.Right)
	if max < l+r+1 {
		max = l + r + 1
	}
	if l > r {
		return l + 1
	}
	return r + 1
}
