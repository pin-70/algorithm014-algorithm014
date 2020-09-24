package Week_06

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid == nil || len(obstacleGrid) == 0 { //异常参数判断
		return 0
	}

	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m) //构建m*n的二维数组
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m && obstacleGrid[i][0] == 0; i++ { //初始化。当[1,0,0]这种情况时，应返回0才对，因为前面被障碍物阻挡了
		dp[i][0] = 1
	}
	for j := 0; j < n && obstacleGrid[0][j] == 0; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = 0
			}
		}
	}
	return dp[m-1][n-1]
}
