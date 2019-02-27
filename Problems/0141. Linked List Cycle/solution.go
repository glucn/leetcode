package solution

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Runtime 4 ms
// Memory 3.8 MB

func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {

		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false

}
