package solution

// Runtime 12 ms
// Memory 6 MB

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	_, t := dfs(root)
	return t
}

func dfs(n *TreeNode) (sum int, tilt int) {
	if n == nil {
		return 0, 0
	}
	sumLeft, tiltLeft := dfs(n.Left)
	sumRight, tiltRight := dfs(n.Right)

	if sumLeft < sumRight {
		return sumLeft + sumRight + n.Val, tiltLeft + tiltRight + (sumRight - sumLeft)
	}
	return sumLeft + sumRight + n.Val, tiltLeft + tiltRight + (sumLeft - sumRight)
}
