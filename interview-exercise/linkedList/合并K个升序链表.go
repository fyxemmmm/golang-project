package linkedList

import "container/heap"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	h := nodeHeap{}
	heap.Init(&h)

	for _, list := range lists {
		if list != nil {
			heap.Push(&h, list)
		}
	}

	dummy := &ListNode{}
	list := dummy
	for len(h) > 0 {
		l := heap.Pop(&h).(*ListNode)
		list.Next = l
		list = list.Next
		if l.Next != nil {
			heap.Push(&h, l.Next)
		}
	}

	return dummy.Next
}


type nodeHeap []*ListNode

func (this nodeHeap) Less(i, j int) bool {
	return this[i].Val < this[j].Val
}

func (this nodeHeap) Len() int {
	return len(this)
}

func (this nodeHeap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *nodeHeap) Push(x interface{}) {
	*this = append(*this, x.(*ListNode))
}

func (this *nodeHeap) Pop() interface{} {
	res := (*this)[len(*this)-1]
	*this = (*this)[:len(*this)-1]
	return res
}
