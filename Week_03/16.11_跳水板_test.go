package Week_03

import (
	"fmt"
	"testing"
)

// 传统的方式求解。犯的问题有：1.没有判断k=0。2.没有判断shorter=longer。3.要求从小到大输出，没有细读题目要求。
func divingBoard(shorter int, longer int, k int) []int {
	if k == 0 {
		return nil
	}
	if shorter == longer {
		return []int{shorter * k}
	}
	res := make([]int, k+1)
	for i := 0; i <= k; i++ {
		res[i] = longer*i + shorter*(k-i)
	}
	return res
}

// 等差数列求解，参考https://leetcode-cn.com/problems/diving-board-lcci/solution/li-yong-zui-da-zui-xiao-chang-du-an-zhao-gu-ding-b/
// shorter*k为最小值，longer*k为最大值；而longer-shorter则为差值。
func divingBoard2(shorter int, longer int, k int) []int {
	if k == 0 {
		return nil
	}
	if shorter == longer {
		return []int{shorter * k}
	}
	res := make([]int, k+1)
	for i := 0; i <= k; i++ {
		res[i] = shorter*k + (longer-shorter)*(i) // k个shorter是最小值，每次增加longer-shorter的差距
	}
	return res
}

func Test_divingBoard(t *testing.T) {
	fmt.Println(divingBoard2(1, 2, 3))
}
