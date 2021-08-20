package dp

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	strLen := 1
	start := 0
	dp := make([][]bool, len(s))

	for right := 0; right < len(s); right ++ {
		dp[right] = make([]bool, len(s))
		dp[right][right] = true

		for left := 0; left < right; left ++ {
			if s[left] == s[right] && (right - left <= 2 || dp[left+1][right-1]) {
				dp[left][right] = true

				l := right-left+1
				if l > strLen {
					strLen = l
					start = left
				}
			}
		}
	}
	return s[start: start+strLen]
}
