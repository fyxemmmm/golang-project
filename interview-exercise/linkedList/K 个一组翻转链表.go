package linkedList


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	pre, end := dummy, dummy

	for end.Next != nil {
		for i := 0; i < k && end != nil; i ++ {
			end = end.Next
		}

		if end == nil {
			return dummy.Next
		}

		next := end.Next
		end.Next = nil
		start := pre.Next
		pre.Next = reverse(start)
		start.Next = next

		pre, end = start, start

	}

	return dummy.Next

}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}