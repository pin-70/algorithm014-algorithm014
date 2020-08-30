package Week_03

func getKthMagicNumber(k int) int {
	dp := make([]int, k+1)
	dp[0] = 1
	p3, p5, p7 := 0, 0, 0
	for i := 1; i < k; i++ {
		dp[i] = min(dp[p3]*3, dp[p5]*5, dp[p7]*7)
		if dp[i] == dp[p3]*3 {
			p3++
		}
		if dp[i] == dp[p5]*5 {
			p5++
		}
		if dp[i] == dp[p7]*7 {
			p7++
		}
	}
	return dp[k-1]
}

func min(a, b, c int) int {
	min := 0
	if a < b {
		min = a
	} else {
		min = b
	}
	if min < c {
		return min
	}
	return c
}
