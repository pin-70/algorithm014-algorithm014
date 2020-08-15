package Week_01

func rotate(nums []int, k int) {
	if nums == nil || len(nums) == 0 || k == len(nums) || k == 0 {
		return
	}
	for i := 0; i < k; i++ { //循环k次
		tmp := nums[len(nums)-1] //每次将最后的元素存到中间变量
		for j := len(nums) - 1; j >= 1; j-- { //依次将元素向后移动1位
			nums[j] = nums[j-1]
		}
		nums[0] = tmp
	}
}
