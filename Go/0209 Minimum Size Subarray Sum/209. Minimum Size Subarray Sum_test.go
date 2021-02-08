package main

import "testing"

type testForm struct {
	target int
	nums   []int
	output int
}

func test(data []testForm, t *testing.T) {
	for _, e := range data {
		got := minSubArrayLen(e.target, e.nums)
		if got != e.output {
			t.Errorf("at input %d %d got %d expected %d\n", e.nums, e.target, got, e.output)
		}
	}
}

func Test(t *testing.T) {
	data := []testForm{
		{
			target: 7,
			nums:   []int{2, 3, 1, 2, 4, 3},
			output: 2,
		},
		{
			target: 4,
			nums:   []int{1, 4, 4},
			output: 1,
		},
		{
			target: 11,
			nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			output: 0,
		},
	}
	test(data, t)
}
