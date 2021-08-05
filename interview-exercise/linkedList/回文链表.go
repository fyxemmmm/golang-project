package linkedList

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	arr := []*ListNode{}
	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}

	for i := 0; i < len(arr)/2; i ++ {
		if arr[i].Val != arr[len(arr)-i-1].Val {
			return false
		}
	}
	return true
}
