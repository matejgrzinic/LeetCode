// https://leetcode.com/problems/merge-k-sorted-lists/submissions/

// Runtime: 4 ms, faster than 99.78% of Go online submissions for Merge k Sorted Lists.
// Memory Usage: 5.9 MB, less than 52.94% of Go online submissions for Merge k Sorted Lists.

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	lists := []*ListNode{initLinkedList([]int{1, 4, 5}), initLinkedList([]int{1, 3, 4}), initLinkedList([]int{2, 6})}

	for index := range lists {
		printLinkedList(lists[index])
	}
	fmt.Println()
	printLinkedList(mergeKLists(lists))
	//printLinkedList(mergeTwoLists(l1, l2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	elements := []int{}

	for index := range lists {
		for lists[index] != nil {
			elements = append(elements, lists[index].Val)
			lists[index] = lists[index].Next
		}
	}

	head := &ListNode{}
	last := head
	sort.Ints(elements)

	for _, value := range elements {
		e := &ListNode{
			Val: value,
		}
		last.Next = e
		last = e
	}
	return head.Next
}

func mergeKLists2(lists []*ListNode) *ListNode {
	head := &ListNode{}
	last := head

	ptrs := []*ListNode{}
	for index := range lists {
		ptrs = append(ptrs, lists[index])
	}

	for {
		value := math.MaxInt32
		minIndex := -1
		for index := 0; index < len(ptrs); index++ {
			if ptrs[index] != nil {
				if ptrs[index].Val < value {
					value = ptrs[index].Val
					minIndex = index
				}
			} else {
				removeNil(&ptrs, index)
				index--
			}
		}
		if minIndex == -1 {
			break
		}

		element := &ListNode{
			Val: value,
		}
		last.Next = element
		last = element
		ptrs[minIndex] = ptrs[minIndex].Next
	}
	return head.Next
}

func removeNil(s *[]*ListNode, i int) {
	(*s)[i] = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
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
