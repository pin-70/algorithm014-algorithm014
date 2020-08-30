package Week_03

import (
	"fmt"
	"testing"
)

func climbStairs(n int) int {
	if n == 1 || n == 2 || n == 0 {
		return n
	}
	s1, s2, sum := 1, 2, 0
	for i := 3; i <= n; i++ {
		sum = s1 + s2
		s1 = s2
		s2 = sum
	}
	return sum
}

// 力扣中运行报错超出时间限制
func climbStairs2(n int) int {
	if n == 1 || n == 2 || n == 0 {
		return n
	}
	res := 0
	if n >= 3 {
		res = climbStairs(n-1) + climbStairs(n-2)
		n--
	}
	return res
}

// 力扣中运行报错超出时间限制
func climbStairs3(n int) int {
	if n == 1 || n == 2 || n == 0 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

func Test_ClimbStairs(t *testing.T) {
	fmt.Println(climbStairs3(44))
	fmt.Println(climbStairs2(44))
	fmt.Println(climbStairs(44))
}
