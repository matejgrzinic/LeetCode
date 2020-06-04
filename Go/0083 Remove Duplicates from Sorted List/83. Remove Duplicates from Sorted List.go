package main

import (
	"./test0083"
	. "./test0083"
)

func main() {
	test0083.Test(deleteDuplicates)
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	t, l := head, head
	t = t.Next
	for {
		if t == nil {
			l.Next = nil
			return head
		}
		if t.Val != l.Val {
			l.Next = t
			l = t
		}
		t = t.Next
	}
}
