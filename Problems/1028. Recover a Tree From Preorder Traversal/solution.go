package solution

// Runtime 16 ms
// Memory 8.9 MB

func recoverFromPreorder(S string) *TreeNode {
	var stack []*TreeNode

	for i := 0; i < len(S); i++ {
		level := 0
		for i < len(S) && S[i] == '-' {
			level++
			i++
		}

		num := 0
		for i < len(S) && S[i] != '-' {
			num = num*10 + int(S[i]) - int('0')
			i++
		}

		i--

		stack = stack[:level]

		newNode := &TreeNode{Val: num}

		if len(stack) > 0 {
			if stack[len(stack)-1].Left == nil {
				stack[len(stack)-1].Left = newNode
			} else {
				stack[len(stack)-1].Right = newNode
			}
		}

		stack = append(stack, newNode)

	}

	return stack[0]

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
