package Week_03

// 未完成
func findSubsequences(nums []int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < len(nums)-1; i++ { //依次循环len-1个数字
		for j := 2; j <= len(nums); j++ { //递增子序列的长度依次为2,3,4...
			for k := i + 1; k+j < len(nums); k++ {
				tmp := append([]int{nums[i]}, nums[k:k+j]...)
				res = append(res, tmp)
			}
		}
	}
	return res
}
