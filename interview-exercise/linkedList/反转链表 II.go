package linkedList


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	pre, end := dummy, dummy

	for i := 0; i < left-1; i ++ {
		pre = pre.Next
	}

	for i := 0; i < right; i ++ {
		end = end.Next
	}

	next := end.Next
	end.Next = nil
	start := pre.Next
	pre.Next = reverseV(start)
	start.Next = next

	return dummy.Next
}

func reverseV(head *ListNode) *ListNode {
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