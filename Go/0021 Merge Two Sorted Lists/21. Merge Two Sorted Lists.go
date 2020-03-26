package main

import (
	"fmt"
	"math"
)

func main() {
	l1 := initLinkedList([]int{1, 2, 4})
	l2 := initLinkedList([]int{1, 3, 4})
	printLinkedList(l1)
	printLinkedList(l2)
	fmt.Println()
	printLinkedList(mergeTwoLists(l1, l2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	last := head
	first, second := l1, l2
	for first != nil || second != nil {
		val := math.MaxInt32
		if first != nil {
			val = first.Val
		}
		if second != nil {
			if second.Val < val {
				val = second.Val
				second = second.Next
			} else {
				first = first.Next
			}
		} else {
			first = first.Next
		}
		element := &ListNode{
			Val: val,
		}
		last.Next = element
		last = element
	}
	return head.Next
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
