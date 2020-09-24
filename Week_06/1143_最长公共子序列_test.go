package Week_06

import (
	"fmt"
	"testing"
)

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if text1[i] == text2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m-1][n-1]
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Test_a2(t *testing.T) {
	t1 := "aaa"
	t2 := "aaaaaa"
	fmt.Println(longestCommonSubsequence(t1, t2))
}
