package test0112

import (
	"fmt"
	"strconv"
	"time"
)

// TreeNode is just a structure
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type testFunc func(*TreeNode, int) bool

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

	equals(BinaryTree([]string{"-2", "nil", "-3"}), -5, true)
	equals(BinaryTree([]string{"5", "4", "8", "11", "nil", "13", "4", "7", "2", "nil", "nil", "nil", "nil", "nil", "1"}), 22, true)

	duration := time.Since(start)
	fmt.Println("Ran", numTests, "tests, ", failed, "failed")
	fmt.Println("Tests took:", duration.Milliseconds(), "ms")
}

// BinaryTree initializes a new binary tree
func BinaryTree(nums []string) *TreeNode {
	if len(nums) < 1 {
		return nil
	}

	value, err := strconv.Atoi(nums[0])

	if err != nil {
		return nil
	}

	root := &TreeNode{
		Val: value,
	}

	if len(nums) > 1 {
		l := connectToChild(nums, 1, root)
		if l != nil {
			root.Left = l
		}
	}
	if len(nums) > 2 {
		r := connectToChild(nums, 2, root)
		if r != nil {
			root.Right = r
		}
	}

	return root
}

func connectToChild(nums []string, index int, parent *TreeNode) *TreeNode {
	value, err := strconv.Atoi(nums[index])
	if err != nil {
		return nil
	}
	tmp := &TreeNode{
		Val: value,
	}

	if index*2+1 < len(nums) {
		l := connectToChild(nums, index*2+1, tmp)
		if l != nil {
			tmp.Left = l
		}
	}
	if index*2+2 < len(nums) {
		r := connectToChild(nums, index*2+2, tmp)
		if r != nil {
			tmp.Right = r
		}
	}

	return tmp
}

func equals(p *TreeNode, sum int, expected bool) {
	numTests++
	got := testFunction(p, sum)
	if got != expected {
		fmt.Println(getElements(p), sum, "got:", got, "expected:", expected)
		failed++
	}
}

func getElements(head *TreeNode) []int {
	elements := []int{}
	if head.Left != nil {
		elements = getElements(head.Left)
	}
	elements = append(elements, head.Val)
	if head.Right != nil {
		elementsRight := getElements(head.Right)
		elements = append(elements, elementsRight...)
	}
	return elements
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
