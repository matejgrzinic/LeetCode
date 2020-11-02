package main

import "testing"

type testForm struct {
	nums   []int
	output int
}

func test(data []testForm, t *testing.T) {
	for _, e := range data {
		got := getMaxLen(e.nums)
		if got != e.output {
			t.Errorf("at input %d got %d expected %d\n", e.nums, got, e.output)
		}
	}
}

func Test(t *testing.T) {
	data := []testForm{
		{
			nums:   []int{1, -2, -3, 4},
			output: 4,
		},
		{
			nums:   []int{0, 1, -2, -3, -4},
			output: 3,
		},
		{
			nums:   []int{-1, -2, -3, 0, 1},
			output: 2,
		},
		{
			nums:   []int{-1, 2},
			output: 1,
		},
		{
			nums:   []int{1, 2, 3, 5, -6, 4, 0, 10},
			output: 4,
		},
		{
			nums:   []int{-1},
			output: 0,
		},
		{
			nums:   []int{},
			output: 0,
		},
		{
			nums:   []int{0, 0, 0, 0},
			output: 0,
		},
		{
			nums:   []int{70, -18, 75, -72, -69, -84, 64, -65, 0, -82, 62, 54, -63, -85, 53, -60, -59, 29, 32, 59, -54, -29, -45, 0, -10, 22, 42, -37, -16, 0, -7, -76, -34, 37, -10, 2, -59, -24, 85, 45, -81, 56, 86},
			output: 14,
		},
	}
	test(data, t)
}
