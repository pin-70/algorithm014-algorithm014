package Week_03

import "math"

func reverse(x int) int {
	res := 0
	for x != 0 {
		n := x % 10
		res = res*10 + n
		x = x / 10
	}
	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}
	return res
}
