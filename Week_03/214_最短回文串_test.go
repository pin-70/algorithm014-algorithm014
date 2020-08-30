package Week_03

import (
	"fmt"
	"testing"
)

func Test_huiwen(t *testing.T) {
	fmt.Println(shortestPalindrome("abbacd"))
}

func shortestPalindrome(s string) string {
	if isHuiwen(s) {
		return s
	}

	index, mid := 0, len(s)/2-1
	res := s + s
	for index < mid {
		tmp := ""
		for r := len(s) - 1; r > 2*mid; r-- {
			tmp = tmp + string(s[r])
		}
		if isHuiwen(tmp + s) {
			if len(tmp+s) < len(res) {
				res = tmp + s
			}
		}

		index++
	}
	return res
}

func isHuiwen(s string) bool {
	l := len(s)
	if l == 1 {
		return true
	}
	i, j := 0, l-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
