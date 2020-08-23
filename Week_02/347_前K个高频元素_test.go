package Week_02

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	res := make([]int, 0)
	for k != 0 {
		temp := 0
		max := 0
		for index, v := range m {
			if v >= max {
				max = v
				temp = index
			}
		}
		m[temp] = 0
		res = append(res, temp)
		k--
	}
	return res
}
