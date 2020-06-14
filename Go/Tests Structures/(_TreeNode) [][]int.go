package test####

import (
	"fmt"
	"time"
)


// TreeNode is just a structure
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

type testFunc func(*TreeNode) [][]int

var numTests int
var failed int
var testFunction testFunc

// Test function runs over all tests and prints unsuccessful tests
func Test(fn testFunc) {
	numTests, failed = 0, 0
	testFunction = fn
	start := time.Now()

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

func equals(p *TreeNode, expected [][]int]) {
	numTests++
	got := testFunction(p)
	if !sliceEquals(got, expected) {
		fmt.Println(getElements(p), "got:", got, "expected:", expected)
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


func sliceEquals(slice1 [][]int, slice2 [][]int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		if len(slice1[i]) != len(slice2[i]){
			return false
		}
		for j := 0; j < len(slice1[i]); j++ {
			if slice1[i][j] != slice2[i][j] {
				return false
			}
		}		
	}
	return true
}
