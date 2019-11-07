package solution

// Runtime 0 ms, faster than 100%
// Memory 2.6 MB, less than 100%

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	paths := dfs(root)
	sum := 0
	for _, p := range paths {
		sum += p.sum
	}
	return sum
}

func dfs(node *TreeNode) []path {
	if node == nil {
		return nil
	}
	if node.Left == nil && node.Right == nil {
		return []path{{depth: 1, sum: node.Val}}
	}
	leftPaths := dfs(node.Left)
	rightPaths := dfs(node.Right)

	var result []path

	for _, p := range leftPaths {
		result = append(result, path{depth: p.depth + 1, sum: pow10(p.depth)*node.Val + p.sum})
	}
	for _, p := range rightPaths {
		result = append(result, path{depth: p.depth + 1, sum: pow10(p.depth)*node.Val + p.sum})
	}
	return result
}

type path struct {
	depth int
	sum   int
}

var dpPow10 = []int{1}

func pow10(n int) int {
	if n < len(dpPow10) {
		return dpPow10[n]
	}
	p := dpPow10[len(dpPow10)-1]
	for i := len(dpPow10); i <= n; i++ {
		p = p * 10
		dpPow10 = append(dpPow10, p)
	}
	return dpPow10[n]
}
