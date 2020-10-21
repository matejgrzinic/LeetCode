package main

import "testing"

type testForm struct {
	A      []int
	output bool
}

func test(data []testForm, t *testing.T) {
	for _, e := range data {
		got := canReorderDoubled(e.A)
		if got != e.output {
			t.Errorf("at input %d got %t expected %t\n", e.A, got, e.output)
		}
	}
}

func Test(t *testing.T) {
	data := []testForm{
		{
			A:      []int{3, 1, 3, 6},
			output: false,
		},
		{
			A:      []int{2, 1, 2, 6},
			output: false,
		},
		{
			A:      []int{4, -2, 2, -4},
			output: true,
		},
		{
			A:      []int{1, 2, 4, 16, 8, 4},
			output: false,
		},
		{
			A:      []int{2, 4, 8, 1},
			output: true,
		},
		{
			A:      []int{-5, -2},
			output: false,
		},
		{
			A:      []int{1, 2, 1, -8, 8, -4, 4, -4, 2, -2},
			output: true,
		},
	}
	test(data, t)
}
