package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

// Runtime 20 ms
// Memory 5.2 MB

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	val := l1.Val + l2.Val
	if val < 10 {
		return &ListNode{
			Val:  val,
			Next: addTwoNumbers(l1.Next, l2.Next),
		}
	}

	return &ListNode{
		Val:  val - 10,
		Next: addTwoNumbers(addTwoNumbers(l1.Next, &ListNode{Val: 1, Next: nil}), l2.Next),
	}

}
