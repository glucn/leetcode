package solution

// Runtime 0 ms
// Memory 3.2 MB

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(node *TreeNode, path int) int {
	newPath := path<<1 + node.Val
	if node.Left == nil && node.Right == nil {
		return newPath
	}
	sum := 0
	if node.Left != nil {
		sum += dfs(node.Left, newPath)
	}
	if node.Right != nil {
		sum += dfs(node.Right, newPath)
	}

	return sum
}
