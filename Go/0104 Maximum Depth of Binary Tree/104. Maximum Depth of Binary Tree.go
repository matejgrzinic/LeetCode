package main

import (
	"./test0104"
	. "./test0104"
)

func main() {
	test0104.Test(maxDepth)
}

func maxDepth(root *TreeNode) int {
	l, r := 0, 0

	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	l = 1 + maxDepth(root.Left)
	r = 1 + maxDepth(root.Right)

	if l > r {
		return l
	}
	return r
}
