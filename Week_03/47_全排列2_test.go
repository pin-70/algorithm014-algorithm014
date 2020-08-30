package Week_03

import (
	"fmt"
	"sort"
	"testing"
)

var ures = [][]int{}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums) // 对数组进行排序

	var ur []int
	backtracku(nums, ur)

	return ures
}

func backtracku(nums []int, ur []int) {
	if len(ur) == len(nums) {
		ures = append(ures, ur)
		return
	}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] && inSliceU(ur, nums[i-1]) {
			continue
		}
		ur = append(ur, nums[i])
		backtracku(nums, ur)
		ur = ur[:len(ur)-1]
	}
}

func inSliceU(s []int, e int) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func Test_a03(t *testing.T) {
	nums := []int{1, 1, 2}
	fmt.Println(permuteUnique(nums))
}
