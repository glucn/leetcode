package solution

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Runtime 12 ms, faster than 67.19%
// Memory 7.9 MB, less than 100%

func findFrequentTreeSum(root *TreeNode) []int {
	count := make(map[int]int)

	dfs(root, count)

	max := 0
	var result []int
	for k, v := range count {
		if v > max {
			max = v
			result = []int{k}
		} else if v == max {
			result = append(result, k)
		}
	}
	return result
}

func dfs(node *TreeNode, count map[int]int) int {
	if node == nil {
		return 0
	}

	leftSum := dfs(node.Left, count)
	rightSum := dfs(node.Right, count)
	sum := leftSum + rightSum + node.Val
	count[sum] = count[sum] + 1
	return sum
}
