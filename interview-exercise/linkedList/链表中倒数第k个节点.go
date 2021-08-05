package linkedList


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
	arr := []*ListNode{}
	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}

	return arr[len(arr)-k]
}