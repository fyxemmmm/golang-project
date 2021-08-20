package dp

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i ++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		}else {
			dp[i] = nums[i]
		}
	}

	res := dp[0]
	for i := 1; i < len(dp); i ++ {
		if dp[i] > res {
			res = dp[i]
		}
	}

	return res
}
