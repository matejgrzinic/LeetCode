package test0002

import (
	"fmt"
	"time"
)

// ListNode is just a structure
type ListNode struct {
	Val  int
	Next *ListNode
}

type testFunc func(*ListNode, *ListNode) *ListNode

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

	equals(LinkedList([]int{5}), LinkedList([]int{2}), LinkedList([]int{7}))
	equals(LinkedList([]int{5}), LinkedList([]int{5}), LinkedList([]int{0, 1}))
	equals(LinkedList([]int{1}), LinkedList([]int{0}), LinkedList([]int{1}))
	equals(LinkedList([]int{5, 5, 5}), LinkedList([]int{1}), LinkedList([]int{6, 5, 5}))

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

// LinkedList initializes a new Linked List
func LinkedList(nums []int) *ListNode {
	if len(nums) < 1 {
		return nil
	}
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

// PrintLinkedList prints all elements of linked list
func PrintLinkedList(head *ListNode) {
	for tmp := head; tmp != nil; tmp = tmp.Next {
		fmt.Print(tmp.Val, " ")
	}
	fmt.Println()
}

func equals(n *ListNode, m *ListNode, expected *ListNode) {
	numTests++
	got := testFunction(n, m)
	if !linkedListEquals(got, expected) {
		fmt.Println(getElements(n), getElements(m), "got:", getElements(got), "expected:", getElements(expected))
		failed++
	}
}

func getElements(head *ListNode) []int {
	elements := []int{}
	for tmp := head; tmp != nil; tmp = tmp.Next {
		elements = append(elements, tmp.Val)
	}
	return elements
}

func linkedListEquals(l1 *ListNode, l2 *ListNode) bool {
	return sliceEquals(getElements(l1), getElements(l2))
}

func sliceEquals(slice1 []int, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
