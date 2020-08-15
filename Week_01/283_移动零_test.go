package Week_01

// 使用快慢指针。i表示慢指针，用来跟踪不为0的数；j为快指针，来遍历数组
// 当下标j对应的值为非0，且跟下标i不相等时，需要将j的值赋值给下标i，同时将下标j的值置为0
func moveZeroes(nums []int) {
	for i, j := 0, 0; j < len(nums); j++ {
		if nums[j] != 0 {
			if j != i {
				nums[i] = nums[j]
				nums[j] = 0
			}
			i++
		}
	}
}
