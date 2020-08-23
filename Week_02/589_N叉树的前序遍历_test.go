package Week_02

func preorder(root *Node) []int {
	res := make([]int, 0)
	if root != nil {
		res = append(res, root.Val)
		for _, n := range root.Children {
			res = append(res, preorder(n)...)
		}
	}
	return res
}
