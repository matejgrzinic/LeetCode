package main

import (
	"./test0002"
	. "./test0002"
)

func main() {
	test0002.Test(addTwoNumbers)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	first, second := l1, l2

	head := &ListNode{}
	last := head

	carry := 0
	done := false

	for !done || carry == 1 {
		sum := 0 + carry
		carry = 0
		done = true

		if first != nil {
			sum += first.Val
			first = first.Next
			if first != nil {
				done = false
			}
		}
		if second != nil {
			sum += second.Val
			second = second.Next
			if second != nil {
				done = false
			}
		}

		if sum > 9 {
			sum -= 10
			carry = 1
		}

		node := &ListNode{
			Val: sum,
		}

		last.Next = node
		last = node

	}

	return head.Next
}
