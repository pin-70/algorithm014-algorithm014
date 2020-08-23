package Week_02

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	queue := make([]*Node, 0, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		l := len(queue)
		level := make([]int, 0, 0)
		for i := 0; i < l; i += 1 {
			level = append(level, queue[i].Val)
			if queue[i].Children != nil {
				for _, curNode := range queue[i].Children {
					if curNode != nil {
						queue = append(queue, curNode)
					}
				}
			}
		}
		res = append(res, level)
		queue = append(queue[l:])
	}
	return res
}
