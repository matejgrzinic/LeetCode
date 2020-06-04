package main

import (
	"./test0061"
	. "./test0061"
)

func main() {
	test0061.Test(rotateRight)
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 {
		return head
	}
	fast := head
	size := 0
	ok := true
	for i := 0; i < k; i++ {
		if fast != nil {
			fast = fast.Next
			size++
		} else {
			ok = false
			break
		}
	}
	if fast == nil && size == k || size == 0 {
		return head
	}
	if ok {
		slow := head
		for fast.Next != nil {
			fast = fast.Next
			slow = slow.Next
			size++
		}
		fast.Next = head
		head = slow.Next
		slow.Next = nil
	} else {
		head = rotateRight(head, k%size)
	}
	return head
}
