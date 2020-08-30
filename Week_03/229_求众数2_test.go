package Week_03

import (
	"testing"
)

func majorityElement2(nums []int) []int {
	res := make([]int, 0)
	mp := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		mp[nums[i]]++
	}
	for k, v := range mp {
		if v > len(nums)/3 {
			res = append(res, k)
		}
	}
	return res
}

func Test_aaaa011(t *testing.T) {
	majorityElement([]int{6, 6, 6, 7, 7})
}

func majorityElement(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	var num, count [2]int

	for _, v := range nums {
		switch {
		case v == num[0]:
			count[0]++
		case v == num[1]:
			count[1]++
		case count[0] == 0:
			num[0] = v
			count[0]++
		case count[1] == 0:
			num[1] = v
			count[1]++
		default:
			count[0]--
			count[1]--
		}
	}

	for k, _ := range count {
		count[k] = 0
	}

	for _, v := range nums {
		switch {
		case num[0] == v:
			count[0]++
		case num[1] == v:
			count[1]++
		default:
			continue
		}
	}

	res := make([]int, 0)
	for k, v := range count {
		if v > len(nums)/3 {
			res = append(res, num[k])
		}
	}
	return res
}
