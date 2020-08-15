package Week_01

func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	i := 0
	j := 1
	for ; j < len(nums); j++ {
		if nums[j] != nums[i] {
			if j-i > 1 {
				nums[i+1] = nums[j]
			}
			i++
		}
	}
	return i + 1
}
