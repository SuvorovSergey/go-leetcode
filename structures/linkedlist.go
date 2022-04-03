package structures

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func Nums2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := &ListNode{}
	head := l
	for _, v := range nums {				
		head.Next = &ListNode{Val: v}
		head = head.Next
	}

	return l.Next
}

func (l *ListNode) Display() {
	head := l
	for head != nil {
		fmt.Printf("%v -> ", head.Val)
		head = head.Next
	}
}

func List2Nums(l *ListNode) []int {
	res := []int{}
	current := l

	for current != nil {
		res = append(res, current.Val)
		current = current.Next
	}

	return res
}
