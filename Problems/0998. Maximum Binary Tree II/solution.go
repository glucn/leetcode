package solution

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Runtime 4 ms
// Memory 3.6 MB

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	list := getList(root)
	newList := append(list, val)

	return buildTree(newList, 0, len(newList))
}

func getList(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return append(append(getList(root.Left), root.Val), getList(root.Right)...)
}

func buildTree(nums []int, l, r int) *TreeNode {
	if l == r {
		return nil
	}
	maxIndex := max(nums, l, r)
	return &TreeNode{
		Val:   nums[maxIndex],
		Left:  buildTree(nums, l, maxIndex),
		Right: buildTree(nums, maxIndex+1, r),
	}
}

func max(nums []int, l, r int) int {
	maxIndex := l
	for i := l; i < r; i++ {
		if nums[maxIndex] < nums[i] {
			maxIndex = i
		}
	}
	return maxIndex
}
