package _872__Leaf_Similar_Trees

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Runtime 0 ms, faster than 100%
// Memory 2.6 MB, less than 50%

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leafs1 := dfs(root1)
	leafs2 := dfs(root2)

	if len(leafs1) != len(leafs2) {
		return false
	}
	for i := range leafs1 {
		if leafs1[i] != leafs2[i] {
			return false
		}
	}
	return true
}

func dfs(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}
	if node.Left == nil && node.Right == nil {
		return []int{node.Val}
	}
	left := dfs(node.Left)
	right := dfs(node.Right)
	return append(left, right...)
}
