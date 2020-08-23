package Week_02

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		a := target - nums[i]
		if _, ok := mp[nums[i]]; !ok {
			mp[nums[i]] = i
		} else {
			return []int{i, nums[i]}
		}
	}
	return nil
}
