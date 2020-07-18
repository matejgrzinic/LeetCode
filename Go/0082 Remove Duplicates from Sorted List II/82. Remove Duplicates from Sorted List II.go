package main

import (
	"./test0082"
	. "./test0082"
)

func main() {
	test0082.Test(deleteDuplicates)
}

func deleteDuplicates(head *ListNode) *ListNode {
	var newHead *ListNode
	var newHeadNext *ListNode

	for tmp := head; tmp != nil; tmp = tmp.Next {
		if tmp.Next != nil && tmp.Val == tmp.Next.Val {
			for tmp.Next != nil && tmp.Val == tmp.Next.Val {
				tmp = tmp.Next
			}
		} else {
			if newHead == nil {
				newHead = &ListNode{Val: tmp.Val}
				newHeadNext = newHead
			} else {
				newHeadNext.Next = &ListNode{Val: tmp.Val}
				newHeadNext = newHeadNext.Next
			}
		}
	}
	if newHead == nil {
		return nil
	}
	return newHead
}
