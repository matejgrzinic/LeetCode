package main

import "testing"

type testForm struct {
	weights []int
	D       int
	output  int
}

func test(data []testForm, t *testing.T) {
	for _, e := range data {
		got := shipWithinDays(e.weights, e.D)
		if got != e.output {
			t.Errorf("at input %d %d got %d expected %d\n", e.weights, e.D, got, e.output)
		}
	}
}

func Test(t *testing.T) {
	data := []testForm{
		{
			weights: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			D:       5,
			output:  15,
		},
		{
			weights: []int{3, 2, 2, 4, 1, 4},
			D:       3,
			output:  6,
		},
		{
			weights: []int{1, 2, 3, 1, 1},
			D:       4,
			output:  3,
		},
		{
			weights: []int{10, 50, 100, 100, 50, 100, 100, 100},
			D:       5,
			output:  160,
		},
		{
			weights: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			D:       1,
			output:  55,
		},
		{
			weights: []int{1, 1, 1, 1, 1},
			D:       1,
			output:  5,
		},
	}
	test(data, t)
}
