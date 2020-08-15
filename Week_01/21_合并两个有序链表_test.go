package Week_01

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var res *ListNode

	if l1.Val < l2.Val { //之前在前面多加了一层for l1!=nil && l2!=nil导致堆栈溢出了
		res = l1
		l1 = l1.Next
		res.Next = mergeTwoLists(l1, l2)
	} else {
		res = l2
		l2 = l2.Next
		res.Next = mergeTwoLists(l2, l1)
	}
}
