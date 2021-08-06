package topk

// 快速排序法
func findKthLargestByQuickSort(nums []int, k int) int {
	length := len(nums)
	left := 0
	right := length - 1
	target := length - k

	for {
		index := partition(nums, left, right)
		if index == target {
			return nums[target]
		}else if index < target {
			left = index + 1
		}else {
			right = index - 1
		}
	}
}



func findKthLargestByHeap(nums []int, k int) int {
	heapSize := len(nums)
	buildHeap(nums, heapSize)

	for i := len(nums)-1; i >= len(nums)-k+1; i -- {
		nums[0] = nums[i]
		heapSize --
		maxHeapify(nums, 0, heapSize)
	}

	return nums[0]
}