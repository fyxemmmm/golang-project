package linkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	first := head.Next
	second := head
	for first != nil {
		if first.Val == second.Val {
			first = first.Next
			second.Next = second.Next.Next
			continue
		}

		first = first.Next
		second = second.Next
	}

	return head
}