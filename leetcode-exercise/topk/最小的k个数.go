package topk


func getLeastNumbersByQuickSort(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	length := len(arr)
	left := 0
	right := length - 1

	for {
		index := partition(arr, left, right)
		if index == k - 1{
			return arr[:k]
		}else if index < k - 1 {
			left = index + 1
		}else {
			right = index - 1
		}
	}
}


func getLeastNumbersByHeap(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}

	buildHeap(arr, k)

	for i := k; i < len(arr); i ++ {
		if arr[i] < arr[0] {
			arr[0] = arr[i]
			maxHeapify(arr, 0, k)
		}
	}
	return arr[:k]
}
