package Week_02

import "math"

func nthUglyNumber(n int) int {
	a, b, c := 0, 0, 0
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		n2 := dp[a] * 2
		n3 := dp[b] * 3
		n5 := dp[c] * 5
		dp[i] = int(math.Min(math.Min(float64(n2), float64(n3)), float64(n5)))
		if dp[i] == n2 {
			a++
		}
		if dp[i] == n3 {
			b++
		}
		if dp[i] == n5 {
			c++
		}
	}
	return dp[n-1]
}
