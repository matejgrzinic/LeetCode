// // https://leetcode.com/problems/add-two-numbers/

// // Runtime: 8 ms, faster than 87.11% of Go online submissions for Add Two Numbers.
// // Memory Usage: 5 MB, less than 12.20% of Go online submissions for Add Two Numbers.

package main

import "fmt"

func main() {
	l1 := initLinkedList([]int{5})
	l2 := initLinkedList([]int{5})
	printLinkedList(l1)
	printLinkedList(l2)
	fmt.Println()
	printLinkedList(addTwoNumbers(l1, l2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func initLinkedList(nums []int) *ListNode {
	head := &ListNode{
		Val: nums[0],
	}

	var last *ListNode

	for index, num := range nums[1:] {
		node := &ListNode{
			Val: num,
		}
		if index == 0 {
			head.Next = node
		} else {
			last.Next = node
		}
		last = node
	}

	return head
}

func printLinkedList(head *ListNode) {
	for tmp := head; tmp != nil; tmp = tmp.Next {
		fmt.Print(tmp.Val, " ")
	}
	fmt.Println()
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
