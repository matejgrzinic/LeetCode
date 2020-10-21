package main

import "testing"

type testForm struct {
	A      []int
	output bool
}

func test(data []testForm, t *testing.T) {
	for _, e := range data {
		got := PredictTheWinner(e.A)
		if got != e.output {
			t.Errorf("at input %d got %t expected %t\n", e.A, got, e.output)
		}
	}
}

func Test(t *testing.T) {
	data := []testForm{
		{
			A:      []int{1, 5, 2},
			output: false,
		},
		{
			A:      []int{1, 5, 233, 7},
			output: true,
		},
		{
			A:      []int{0},
			output: true,
		},
		{
			A:      []int{0, 1},
			output: true,
		},
		{
			A:      []int{1, 1},
			output: true,
		},
	}
	test(data, t)
}
