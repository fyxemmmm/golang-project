package dp

func findLength(nums1 []int, nums2 []int) int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return 0
	}

	dp := make([][]int, len(nums1)+1)
	for i := 0; i < len(dp); i ++ {
		dp[i] = make([]int, len(nums2)+1)
	}

	res := 0

	for i := 1; i < len(dp); i ++ {
		for j := 1; j < len(dp[i]); j ++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				res = max(res, dp[i][j])
			}
		}
	}
	return res
}
