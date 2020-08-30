package Week_03

import "container/list"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 其他思路：1.利用栈的先进后出，pop到切片中来做。2.先反转链表，而后再遍历，直接将值输出到切片。3.利用递归的方式解决

// 利用切片来存储链表中的值。待遍历完毕后，将数组头尾交换。
func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}

	res := make([]int, 0)
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	for i := 0; i < len(res)/2; i++ {
		tmp := res[i]
		res[i] = res[len(res)-1-i]
		res[len(res)-1-i] = tmp
	}
	return res
}

func reversePrint2(head *ListNode) []int {
	stack := list.New()
	res := make([]int, 0)

	for head != nil {
		stack.PushBack(head.Val)
		head = head.Next
	}

	for stack.Len() > 0 {
		e := stack.Back()
		stack.Remove(e)
		res = append(res, e.Value.(int))
	}
	return res
}
