package Week_01

// 借助hashmap来实现
func twoSum2(nums []int, target int) []int {
	tmp := make(map[int]int) //k=value,v=index
	for i := 0; i < len(nums); i++ {
		mid := target - nums[i]
		if _, ok := tmp[mid]; !ok {
			tmp[nums[i]] = i
		}
		return []int{i, tmp[mid]}
	}
	return nil
}

//借助双指针来做，注意判重。当然也可以不判断
//在力扣中执行时发现：加上判重后，时间比不加判重多消耗了44ms
func twoSum3(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
