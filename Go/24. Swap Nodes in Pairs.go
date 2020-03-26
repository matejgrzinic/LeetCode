// https://leetcode.com/problems/swap-nodes-in-pairs/

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Swap Nodes in Pairs.
// Memory Usage: 2.1 MB, less than 66.67% of Go online submissions for Swap Nodes in Pairs.

package main

import (
	"fmt"
)

func main() {
	l1 := initLinkedList([]int{1, 2, 3, 4, 5, 6})
	printLinkedList(swapPairs(l1))
}

func swapPairs(head *ListNode) *ListNode {
	tmp := head
	var last *ListNode
	first := true

	for tmp != nil {
		if tmp.Next == nil {
			return head
		}
		if first {
			head = swapNode(tmp, tmp.Next)
			first = false
		} else {
			last.Next = swapNode(tmp, tmp.Next)
		}
		last = tmp
		tmp = tmp.Next
	}
	return head
}

func swapNode(node1 *ListNode, node2 *ListNode) *ListNode {
	node1.Next = node2.Next
	node2.Next = node1
	return node2
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
