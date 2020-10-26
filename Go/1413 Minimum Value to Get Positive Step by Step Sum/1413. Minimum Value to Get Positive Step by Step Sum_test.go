package main

import "testing"

type testForm struct {
	nums   []int
	output int
}

func test(data []testForm, t *testing.T) {
	for _, e := range data {
		got := minStartValue(e.nums)
		if got != e.output {
			t.Errorf("at input %d got %d expected %d\n", e.nums, got, e.output)
		}
	}
}

func Test(t *testing.T) {
	data := []testForm{
		{
			nums:   []int{-3, 2, -3, 4, 2},
			output: 5,
		},
		{
			nums:   []int{1, 2},
			output: 1,
		},
		{
			nums:   []int{1, -2, -3},
			output: 5,
		},
	}
	test(data, t)
}
