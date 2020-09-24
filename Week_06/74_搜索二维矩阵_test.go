package Week_06

import "testing"

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0]) //m代表行数,n代表列数
	l, r := 0, m*n-1    // 转成一维升序数组后的左右下标初始值

	for l <= r {
		mid := l + (r-l)/2
		i, j := mid/n, mid%n
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}

func Test_aa(t *testing.T) {
	aa := [][]int{{1, 1}}
	searchMatrix(aa, 3)
}
