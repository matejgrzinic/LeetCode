// https://leetcode.com/problems/reverse-nodes-in-k-group/

// Runtime: 4 ms, faster than 93.66% of Go online submissions for Reverse Nodes in k-Group.
// Memory Usage: 4 MB, less than 25.00% of Go online submissions for Reverse Nodes in k-Group.

package main

import (
	"fmt"
)

func main() {
	l1 := initLinkedList([]int{1, 2, 3, 4, 5})
	printLinkedList(reverseKGroup(l1, 8))
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	first := true
	var last *ListNode

	tmp := &ListNode{
		Next: head,
	}

	for tmp != nil {
		group := []*ListNode{}
		for i := 0; tmp != nil && i < k; i++ {
			tmp = tmp.Next
			group = append(group, tmp)
		}
		if tmp == nil {
			return head
		}
		if first {
			head, tmp = swapNodes(group)
			first = false
		} else {
			last.Next, tmp = swapNodes(group)
		}
		last = tmp
	}

	return head
}

func swapNodes(nodes []*ListNode) (*ListNode, *ListNode) {
	n := len(nodes) - 1
	var next *ListNode
	for i := n; i > 0; i-- {
		if i == n {
			next = nodes[i].Next
		}
		nodes[i].Next = nodes[i-1]
	}
	nodes[0].Next = next
	return nodes[n], nodes[0]
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
