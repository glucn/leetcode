package _080__Insufficient_Nodes_in_Root_to_Leaf_Paths

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Runtime 60 ms
// Memory 42.6 MB

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	isSufficient := dfs(root, limit)
	if isSufficient {
		return root
	}
	return nil
}

func dfs(node *TreeNode, limit int) bool {
	if node == nil {
		return false
	}
	if node.Left == nil && node.Right == nil {
		return node.Val >= limit
	}
	isLeftSufficient := dfs(node.Left, limit-node.Val)
	isRightSufficient := dfs(node.Right, limit-node.Val)

	if !isLeftSufficient {
		node.Left = nil
	}
	if !isRightSufficient {
		node.Right = nil
	}
	return isLeftSufficient || isRightSufficient
}
