package Week_03

type TreeNode struct {
	Val   int
	Left  *ListNode
	Right *ListNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}

	//left := lowestCommonAncestor(root.Left, p, q)
	//right := lowestCommonAncestor(root.Right, p, q)
	//if left == nil {
	//	return right
	//}
	//if right == nil {
	//	return left
	//}
	return nil
}
