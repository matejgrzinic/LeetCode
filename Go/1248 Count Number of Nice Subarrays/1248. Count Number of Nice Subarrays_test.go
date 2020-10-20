package main

import "testing"

type testForm struct {
	nums   []int
	k      int
	output int
}

func test(data []testForm, t *testing.T) {
	for _, e := range data {
		got := numberOfSubarrays(e.nums, e.k)
		if got != e.output {
			t.Errorf("at input %d %d got %d expected %d\n", e.nums, e.k, got, e.output)
		}
	}
}

func Test(t *testing.T) {
	data := []testForm{
		{
			nums:   []int{1, 1, 2, 1, 1},
			k:      3,
			output: 2,
		},
		{
			nums:   []int{2, 4, 6},
			k:      1,
			output: 0,
		},
		{
			nums:   []int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2},
			k:      2,
			output: 16,
		},
	}
	test(data, t)
}
