package solution

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Runtime 8 ms, faster than 88.83%
// Memory 6.1 MB, less than 50%

func isBalanced(root *TreeNode) bool {
	res, _ := dfs(root)
	return res
}

func dfs(node *TreeNode) (bool, int) {
	if node == nil {
		return true, 0
	}

	isBalanced, depthLeft := dfs(node.Left)
	if !isBalanced {
		return false, 0
	}
	isBalanced, depthRight := dfs(node.Right)
	if !isBalanced {
		return false, 0
	}
	if abs(depthLeft-depthRight) > 1 {
		return false, 0
	}
	return true, max(depthLeft, depthRight) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
