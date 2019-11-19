package solution

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Runtime 4 ms, faster than 36.96% // TODO: make it faster
// Memory 3 MB, less than 100%

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	resultMap := make(map[int]struct{})
	_ = dfs(root, target, K, resultMap)

	result := make([]int, 0, len(resultMap))
	for k := range resultMap {
		result = append(result, k)
	}
	return result
}

func dfs(node *TreeNode, target *TreeNode, K int, result map[int]struct{}) int {
	if node == nil {
		return -1
	}

	if node == target {
		subDFS(node, K, 0, result)
		return 1
	}

	left := dfs(node.Left, target, K, result)
	right := dfs(node.Right, target, K, result)
	if left != -1 {
		if left == K {
			result[node.Val] = struct{}{}
		}
		subDFS(node.Right, K, left+1, result)
		return left + 1
	}

	if right != -1 {
		if right == K {
			result[node.Val] = struct{}{}
		}
		subDFS(node.Left, K, right+1, result)
		return right + 1
	}

	return -1
}

func subDFS(node *TreeNode, K int, distance int, result map[int]struct{}) {
	if node == nil {
		return
	}

	if K == distance {
		result[node.Val] = struct{}{}
		return
	}

	subDFS(node.Left, K, distance+1, result)
	subDFS(node.Right, K, distance+1, result)
}
