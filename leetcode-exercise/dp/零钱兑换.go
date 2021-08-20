package dp

func coinChange(coins []int, amount int) int {
	dp := make([][]int, len(coins)+1)
	for i := 0; i < len(dp); i ++ {
		dp[i] = make([]int, amount+1)
	}

	for j := 1; j <= amount; j ++{
		dp[0][j] = amount + 1
	}

	for i := 1; i < len(dp); i ++ {
		for j := 0; j <= amount; j ++ {
			if j >= coins[i-1] {
				dp[i][j] = min(dp[i-1][j], dp[i][j-coins[i-1]]+1)
			}else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	if dp[len(coins)][amount] == amount + 1 {
		return -1
	}
	return dp[len(coins)][amount]
}
