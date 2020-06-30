package main

import (
	"./test0113"
	. "./test0113"
)

func main() {
	test0113.Test(pathSum)
}

func pathSum(root *TreeNode, sum int) [][]int {
	result := [][]int{}
	helper(root, sum, []int{}, &result)
	return result
}

func helper(root *TreeNode, sum int, current []int, result *[][]int) {
	if root == nil {
		return
	}
	current = append(current, root.Val)
	sum -= root.Val

	if root.Left == nil && root.Right == nil {
		if sum == 0 {
			copyCurrent := make([]int, len(current))
			copy(copyCurrent, current)
			*result = append(*result, copyCurrent)
		}
		return
	}

	helper(root.Left, sum, current, result)
	helper(root.Right, sum, current, result)
}
