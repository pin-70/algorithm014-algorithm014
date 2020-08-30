package Week_03

import (
	"fmt"
	"sort"
	"testing"
)

func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		first := nums[i]
		for j := i + 1; j < len(nums)-2; j++ {
			second := nums[j]
			tmp := target - first - second
			k, l := j+1, len(nums)-1
			for k < l {
				if tmp-nums[k]-nums[l] < 0 {
					l--
				} else if tmp-nums[k]-nums[l] > 0 {
					k++
				} else {
					res = append(res, []int{first, second, nums[k], nums[l]})
					l--
					k++
				}
			}
		}
	}
	return res
}

func Test_ccc01(t *testing.T) {
	fmt.Println(fourSum([]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0))
}
