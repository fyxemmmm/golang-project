package dp

func change(amount int, coins []int) int {
	dp := make([][]int, len(coins)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1
	}

	for i := 1; i < len(dp); i ++ {
		for j := 0; j <= amount; j ++ {
			if j >= coins[i-1] {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			}else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(coins)][amount]
}