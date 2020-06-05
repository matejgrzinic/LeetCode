package main

import (
	"./test0100"
	. "./test0100"
)

func main() {
	test0100.Test(isSameTree)
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil {
		if q == nil {
			return true
		}
		return false
	} else {
		if q == nil {
			return false
		}
	}

	if p.Val != q.Val {
		return false
	}

	l := isSameTree(p.Left, q.Left)
	if !l {
		return false
	}

	r := isSameTree(p.Right, q.Right)
	if !r {
		return false
	}

	return true
}
