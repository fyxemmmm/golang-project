package array

import "sort"

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		target := -nums[i]
		left := i + 1
		right := len(nums) - 1

		for left < right {
			if nums[left]+nums[right] == target {
				res = append(res, []int{nums[left], nums[right], nums[i]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if nums[left]+nums[right] < target {
				left++
			} else {
				right--
			}
		}
	}

	return res
}

