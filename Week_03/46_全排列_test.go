package Week_03

import (
	"fmt"
	"testing"
)

var res = make([][]int, 0)

func permute(nums []int) [][]int {
	r := []int{}
	backtrack(nums, r)
	return res
}

func backtrack(nums []int, r []int) {
	if len(r) == len(nums) {
		res = append(res, r)
		return
	}
	for i := 0; i < len(nums); i++ {
		if inSlice(r, nums[i]) {
			continue
		}
		r = append(r, nums[i])
		backtrack(nums, r)
		r = r[:len(r)-1]
	}
}

func inSlice(s []int, e int) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func Test_a02(t *testing.T) {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}
