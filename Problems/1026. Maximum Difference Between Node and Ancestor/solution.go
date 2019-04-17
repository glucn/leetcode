package main

// Runtime 4 ms
// Memory 4.2 MB

func maxAncestorDiff(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return traverse(root, 100001, -1)
}

func traverse(node *TreeNode, minVal int, maxVal int) int {
	// fmt.Println("calling", node.Val, minVal, maxVal)

	localMax := max(node.Val-minVal, maxVal-node.Val)

	if node.Left == nil && node.Right == nil {
		// fmt.Println("returning", node.Val, node.Val - minVal, maxVal - node.Val)
		return localMax
	}

	newMin := min(minVal, node.Val)
	newMax := max(maxVal, node.Val)

	left, right := -1, -1
	if node.Left != nil {
		left = traverse(node.Left, newMin, newMax)
	}
	if node.Right != nil {
		right = traverse(node.Right, newMin, newMax)
	}
	// fmt.Println("returning", node.Val, left, right)
	return max(max(left, right), localMax)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
