package solution

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Runtime 24 ms // TODO: 51.98%
// Memory 33.2 MB // TODO: 16.67%

func buildTree(preorder []int, inorder []int) *TreeNode {
	n, _ := build(preorder, inorder)
	return n
}

func build(preorder []int, inorder []int) (n *TreeNode, after []int) {
	// fmt.Println("input", preorder, inorder)

	if len(preorder) == 0 || len(inorder) == 0 {
		return nil, preorder
	}
	v := preorder[0]
	after = preorder[1:]
	n = &TreeNode{
		Val: v,
	}
	var idx int
	for i, n := range inorder {
		if n == v {
			idx = i
			break
		}
	}
	n.Left, after = build(after, inorder[:idx])
	n.Right, after = build(after, inorder[idx+1:])

	// fmt.Println("leaving", after)
	return n, after
}
