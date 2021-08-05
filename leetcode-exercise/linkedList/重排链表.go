package linkedList


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
	list := head
	arr := []*ListNode{}
	for list != nil {
		arr = append(arr, list)
		list = list.Next
	}

	left := 0
	right := len(arr)-1

	for left < right {
		arr[left].Next = arr[right]
		left ++
		if left == right {
			break
		}
		arr[right].Next = arr[left]
		right --
	}

	arr[right].Next = nil
	return
}