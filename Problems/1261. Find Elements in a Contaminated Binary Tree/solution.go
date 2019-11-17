package solution

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */

// Runtime 36 ms, faster than ?%
// Memory 8.2 MB, less than ?%

type FindElements struct {
	exist map[int]struct{}
}

func Constructor(root *TreeNode) FindElements {
	m := make(map[int]struct{})
	dfs(root, 0, m)
	return FindElements{
		exist: m,
	}
}

func dfs(node *TreeNode, val int, m map[int]struct{}) {
	if node == nil {
		return
	}
	node.Val = val
	m[val] = struct{}{}
	dfs(node.Left, 2*val+1, m)
	dfs(node.Right, 2*val+2, m)
}

func (fe *FindElements) Find(target int) bool {
	if _, ok := fe.exist[target]; ok {
		return true
	}
	return false
}
