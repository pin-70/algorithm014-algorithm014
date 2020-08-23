package Week_02

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	res = helper(root)
	return res
}

func helper(r *TreeNode) (res []int) {
	if r != nil {
		if r.Left != nil {
			res = append(res, helper(r.Left)...)
		}
		res = append(res, r.Val)
		if r.Right != nil {
			res = append(res, helper(r.Right)...)
		}
	}
	return res
}
