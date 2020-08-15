package Week_01

func trap(height []int) int {
	l := len(height)
	lefts := make([]int, l)
	rights := make([]int, l)

	for i := 1; i < l-1; i++ { // i表示当前的列，求它的左边 的最大的柱子
		if height[i-1] > lefts[i-1] {
			lefts[i] = height[i-1]
		} else {
			lefts[i] = lefts[i-1]
		}
	}
	for i := l - 2; i >= 0; i-- { // i表示当前的列，求它的右边 的最大的柱子
		if height[i+1] > rights[i+1] {
			rights[i] = height[i+1]
		} else {
			rights[i] = rights[i+1]
		}
	}

	sum := 0
	for i := 0; i < l; i++ {
		min := lefts[i]
		if lefts[i] > rights[i] {
			min = rights[i]
		}
		if min > height[i] {
			sum = sum + (min - height[i])
		}
	}
	return sum
}
