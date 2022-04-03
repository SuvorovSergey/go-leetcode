package addtwonumbers

import (
	"github.com/SuvorovSergey/go-leetcode/structures"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = structures.ListNode

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	list := &ListNode{}
	n1 := 0
	n2 := 0
	carry := 0
	head := list
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}

		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}
		head.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		head = head.Next
		carry = (n1 + n2 + carry) / 10

	}
	return list.Next
}
