package linkedList
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow{
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	for head != slow {
		head = head.Next
		slow = slow.Next
	}

	return head


}