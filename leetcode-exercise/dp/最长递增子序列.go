package dp

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	res := 1
	for right := 0; right < len(nums); right ++ {
		dp[right] = 1
		for left := 0; left < right; left ++ {
			if nums[left] < nums[right] {
				dp[right] = max(dp[right], dp[left]+1)
				if dp[right] > res {
					res = dp[right]
				}
			}
		}
	}

	return res
}
