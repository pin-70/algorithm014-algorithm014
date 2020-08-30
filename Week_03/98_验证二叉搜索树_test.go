package Week_03

import "math"

// 1.得用math.MinInt64，不能用MinInt32
func isValidBST(root *TreeNode2) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode2, lowest, highest int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lowest || root.Val >= highest {
		return false
	}
	return helper(root.Left, lowest, root.Val) && helper(root.Right, root.Val, highest)
}
