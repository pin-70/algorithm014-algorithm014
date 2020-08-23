package Week_02

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	res = helper(root)
	return res
}

func helper2(root *TreeNode) []int {
	r := make([]int, 0)
	if root != nil {
		r = append(r, root.Val)
		if root.Left != nil {
			r = append(r, helper(root.Left)...)
		}
		if root.Right != nil {
			r = append(r, helper(root.Right)...)
		}
	}
	return r
}
