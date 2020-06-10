package main

import (
	"./test0107"
	. "./test0107"
)

func main() {
	test0107.Test(levelOrderBottom)
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{{root.Val}}

	l, r := [][]int{}, [][]int{}
	if root.Left != nil {
		l = levelOrderBottom(root.Left)
	}
	if root.Right != nil {
		r = levelOrderBottom(root.Right)
	}

	new := merge(l, r)
	if len(new) != 0 {
		res = append(new, res...)
	}
	return res
}

func merge(l [][]int, r [][]int) [][]int {
	n, m := len(l), len(r)

	arr := [][]int{}
	for {
		row := []int{}
		added := false
		if n > 0 {
			row = append(row, l[n-1]...)
			added = true
			n--
		}
		if m > 0 {
			row = append(row, r[m-1]...)
			added = true
			m--
		}
		if !added {
			break
		}
		arr = append([][]int{row}, arr...)
	}
	return arr
}
