package Week_03

import (
	"fmt"
	"math"
	"testing"
)

// 下面方法导致超时，用时14.29s
func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	res := float64(1)

	for i := 1; i <= int(math.Abs(float64(n))); i++ {
		res = res * x
	}
	if n < 0 {
		return float64(1) / res
	}
	return res
}

func myPow5(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	res := float64(1)

	for i := 1; i <= int(math.Abs(float64(n))); i++ {
		res = res * x
	}
	if n < 0 {
		return float64(1) / res
	}
	return res
}

func myPow2(x float64, n int) float64 {

	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}

	if n < 0 {
		x = float64(1.0) / x
		n = -n
	}

	if n%2 == 0 {
		return myPow(x*x, n/2)
	} else {
		return myPow(x*x, n/2) * x
	}
	return 0
}

func Test_bb01(t *testing.T) {
	fmt.Println(myPow(0.00001, 2147483647))
}
