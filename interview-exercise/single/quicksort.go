package main

func quickSort(nums []int, low, high int)  {
	if low >= high {
		return
	}
	left, right := low, high
	pivot := nums[left]

	for left < right {
		if left < right && nums[right] >= pivot {
			right --
		}

		if left < right {
			nums[left] = nums[right]
		}

		if left < right && nums[left] < pivot {
			left ++
		}

		if left < right {
			nums[right] = nums[left]
		}

		if left <= right {
			nums[left] = pivot
		}
	}

	quickSort(nums, low, right-1)
	quickSort(nums, right+1, high)
}