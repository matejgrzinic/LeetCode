// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Nth Node From End of List.
// Memory Usage: 2.2 MB, less than 28.57% of Go online submissions for Remove Nth Node From End of List.

package main

import "fmt"

// func main() {
// 	// head := initLinkedList([]int{1, 2, 3, 4, 5})
// 	head := initLinkedList([]int{1})
// 	printLinkedList(head)
// 	head = removeNthFromEnd(head, 1)
// 	printLinkedList(head)
// }

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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow, front := head, head

	i := 0

	for front != nil && i < n {
		front = front.Next
		i++
	}

	if front == nil {
		if i == n {
			return head.Next
		}
		return head
	}

	for front.Next != nil {
		front = front.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return head
}
