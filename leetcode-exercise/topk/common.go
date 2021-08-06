package topk

import "math/rand"

// 快排
func partition(nums []int, left, right int) int {
	randomIndex := rand.Int() % (right-left+1) + left
	nums[right], nums[randomIndex] = nums[randomIndex], nums[right]

	pivot := nums[right]
	pivotIndex := left

	for i := left; i < right; i ++ {
		if nums[i] < pivot {
			nums[i], nums[pivotIndex]= nums[pivotIndex], nums[i]
			pivotIndex ++
		}
	}

	nums[right], nums[pivotIndex] = nums[pivotIndex], nums[right]
	return pivotIndex
}



// 建堆
func buildHeap(nums []int, heapSize int) {
	for i := len(nums)/2; i >= 0; i -- {
		maxHeapify(nums, i, heapSize)
	}
}


func maxHeapify(nums []int, i, heapSize int) {
	left, right, max := i*2+1, i*2+2, i
	if left < heapSize && nums[left] > nums[max] {
		max = left
	}

	if right < heapSize && nums[right] > nums[max] {
		max = right
	}

	if i != max {
		nums[max], nums[i] = nums[i], nums[max]
		maxHeapify(nums, max, heapSize)
	}
}