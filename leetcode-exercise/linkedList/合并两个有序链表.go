package linkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	list := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			list.Next = &ListNode{
				Val: l1.Val,
			}
			l1 = l1.Next
		}else {
			list.Next = &ListNode{
				Val: l2.Val,
			}
			l2 = l2.Next
		}

		list = list.Next
	}

	if l1 != nil {
		list.Next = l1
	}

	if l2 != nil {
		list.Next = l2
	}

	return dummy.Next
}