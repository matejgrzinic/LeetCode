package main

import (
	"./test0101"
	. "./test0101"
)

func main() {
	test0101.Test(isSymmetric)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left == nil || root.Right == nil {
		return false
	}
	return checkChilds(root.Left, root.Right)
}

func checkChilds(left *TreeNode, right *TreeNode) bool {
	if left.Val != right.Val {
		return false
	}
	l, r := true, true
	if left.Left == nil && right.Right == nil {
		l = false
	} else if left.Left == nil || right.Right == nil {
		return false
	}
	if left.Right == nil && right.Left == nil {
		r = false
	} else if left.Right == nil || right.Left == nil {
		return false
	}

	if l {
		if !checkChilds(left.Left, right.Right) {
			return false
		}
	}
	if r {
		if !checkChilds(left.Right, right.Left) {
			return false
		}
	}
	return true
}
