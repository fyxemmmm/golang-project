package binarySearch

// left <= right
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums)-1

	for left <= right {
		mid := (left+right) >> 1
		if nums[mid] == target {
			return mid
		}else if target >= nums[left] && target <= nums[mid] || (nums[mid] < nums[right] && (target < nums[mid] || target > nums[right])) {
			right = mid - 1
		}else {
			left = mid + 1
		}
	}

	return -1
}