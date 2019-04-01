package solution

// Runtime 620 ms
// Memory 8.4 MB

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {
	node := head
	var values []int
	for node != nil {
		values = append(values, node.Val)
		node = node.Next
	}

	result := make([]int, len(values))
	for i, v := range values {
		result[i] = findLarger(values[i:], v)
	}
	return result
}

func findLarger(list []int, val int) int {
	for _, v := range list {
		if v > val {
			return v
		}
	}
	return 0
}
