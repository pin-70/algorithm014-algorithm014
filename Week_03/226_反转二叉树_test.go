package Week_03

type TreeNode2 struct {
	Val   int
	Left  *TreeNode2
	Right *TreeNode2
}

func invertTree(root *TreeNode2) *TreeNode2 {
	if root == nil || (root != nil && root.Left == nil && root.Right == nil) {
		return root
	}
	if root.Left != nil || root.Right != nil {
		root.Left, root.Right = root.Right, root.Left
	}
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

// 去掉不必要的判断条件
func invertTree2(root *TreeNode2) *TreeNode2 {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
