package main

import "testing"

type testForm struct {
	nums   []int
	k      int
	output float64
}

func test(data []testForm, t *testing.T) {
	for _, e := range data {
		got := findMaxAverage(e.nums, e.k)
		if got != e.output {
			t.Errorf("at input %d %d got %f expected %f\n", e.nums, e.k, got, e.output)
		}
	}
}

func Test(t *testing.T) {
	data := []testForm{
		{
			nums:   []int{1, 12, -5, -6, 50, 3},
			k:      4,
			output: 12.75,
		},
	}
	test(data, t)
}
