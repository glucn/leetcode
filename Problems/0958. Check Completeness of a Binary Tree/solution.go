package solution

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Runtime 0 ms, faster than 100%
// Memory 3.6 MB, less than 100%

func isCompleteTree(root *TreeNode) bool {
	queue := []node{{n: root, code: 1}}
	i := 0
	for i < len(queue) {
		current := queue[i]
		i++
		if current.n != nil {
			queue = append(queue, node{n: current.n.Left, code: current.code * 2})
			queue = append(queue, node{n: current.n.Right, code: current.code*2 + 1})
		}
	}
	return queue[len(queue)-1].code == len(queue)
}

type node struct {
	n    *TreeNode
	code int
}
