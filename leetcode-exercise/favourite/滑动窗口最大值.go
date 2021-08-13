package favourite

func maxSlidingWindow(nums []int, k int) []int {
	if k == 1 {
		return nums
	}

	res := []int{}
	queue := []int{0}

	for i := 0; i < len(nums); i++ {
		for i-k >= queue[0] {
			queue = queue[1:]
		}

		for len(queue) > 0 && nums[i] > nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}

		queue = append(queue, i)

		if i >= k-1 {
			res = append(res, nums[queue[0]])
		}
	}
	return res
}
