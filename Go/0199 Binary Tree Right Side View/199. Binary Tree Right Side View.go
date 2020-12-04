package main

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	rightSideElements := []int{}

	traverse(root, 1, &rightSideElements)
	return rightSideElements
}

func traverse(r *TreeNode, d int, rightSideElements *[]int) {
	if r != nil {
		if d > len(*rightSideElements) {
			*rightSideElements = append(*rightSideElements, r.Val)
		}
		traverse(r.Right, d+1, rightSideElements)
		traverse(r.Left, d+1, rightSideElements)
	}
}
