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
	res := [][]int{}
	order(root, 1, &res)
	for i := 0; i < len(res); i++ {
		j := len(res) - 1 - i
		if i < j {
			res[i], res[j] = res[j], res[i]
		} else {
			break
		}
	}
	return res
}

func order(e *TreeNode, depth int, arr *[][]int) {
	if len(*arr) < depth {
		*(arr) = append(*arr, []int{e.Val})
	} else {
		(*arr)[depth-1] = append((*arr)[depth-1], e.Val)
	}

	if (*e).Left != nil {
		order(e.Left, depth+1, arr)
	}
	if (*e).Right != nil {
		order(e.Right, depth+1, arr)
	}
}
