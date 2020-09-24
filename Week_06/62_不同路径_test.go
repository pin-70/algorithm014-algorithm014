package Week_06

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// 使用一维数组来解决
func uniquePaths2(m int, n int) int {
	dp := make([]int, m)
	for i := 0; i < m; i++ {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[j] = dp[j-1] + dp[j]
		}
	}
	return dp[m-1]
}
