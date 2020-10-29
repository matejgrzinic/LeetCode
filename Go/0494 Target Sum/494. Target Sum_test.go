package main

import "testing"

type testForm struct {
	nums   []int
	S      int
	output int
}

func test(data []testForm, t *testing.T) {
	for _, e := range data {
		got := findTargetSumWays(e.nums, e.S)
		if got != e.output {
			t.Errorf("at input %d %d got %d expected %d\n", e.nums, e.S, got, e.output)
		}
	}
}

func Test(t *testing.T) {
	data := []testForm{
		{
			nums:   []int{1, 1, 1, 1, 1},
			S:      3,
			output: 5,
		},
	}
	test(data, t)
}
