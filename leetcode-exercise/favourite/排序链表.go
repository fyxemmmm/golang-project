package favourite

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fast, slow, preSlow := head, head, head

	for fast != nil && fast.Next != nil {
		preSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	preSlow.Next = nil
	l1 := sortList(head)
	l2 := sortList(slow)
	return merge(l1, l2)
}

func merge(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	list := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			list.Next = &ListNode{
				Val: l1.Val,
			}
			l1 = l1.Next
		} else {
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
