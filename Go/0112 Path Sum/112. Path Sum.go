package main

import (
	"./test0112"
	. "./test0112"
)

func main() {
	test0112.Test(hasPathSum)
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	sum -= root.Val

	if root.Left == nil && root.Right == nil {
		if sum == 0 {
			return true
		}
	}
	if root.Left != nil {
		if hasPathSum(root.Left, sum) {
			return true
		}
	}
	if root.Right != nil {
		if hasPathSum(root.Right, sum) {
			return true
		}
	}
	return false
}
