package main

import "./test0084"

func main() {
	test0084.Test(largestRectangleArea)
}

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	result := 0
	stack := []int{0}

	for i := 1; i < len(heights); i++ {
		num := heights[i]

		if num < heights[stack[len(stack)-1]] {
			for len(stack) > 0 && num < heights[stack[len(stack)-1]] {
				tmp := heights[stack[len(stack)-1]]
				stack = stack[:len(stack)-1]
				if len(stack) == 0 {
					if tmp*i > result {
						result = tmp * i
					}
				} else {
					sum := tmp * (i - stack[len(stack)-1] - 1)
					if sum > result {
						result = sum
					}
				}
			}
		}
		stack = append(stack, i)
	}

	i := len(heights)
	for len(stack) > 0 {
		tmp := heights[stack[len(stack)-1]]
		stack = stack[:len(stack)-1]
		if len(stack) == 0 {
			if tmp*i > result {
				result = tmp * i
			}
		} else {
			sum := tmp * (i - stack[len(stack)-1] - 1)
			if sum > result {
				result = sum
			}
		}
	}
	return result
}
