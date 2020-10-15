package main

import "testing"

type testForm struct {
	input  []int
	output int
}

func testAll(data []testForm, t *testing.T) {
	for _, e := range data {
		got := lastStoneWeightII(e.input)
		if got != e.output {
			t.Errorf("at input %d got %d expected %d\n", e.input, got, e.output)
		}
	}
}

func Test1(t *testing.T) {
	data := []testForm{
		{
			input:  []int{2, 7, 4, 1, 8, 1},
			output: 1,
		},
		{
			input:  []int{},
			output: 0,
		},
		{
			input:  []int{1},
			output: 1,
		},
		{
			input:  []int{1, 1},
			output: 0,
		},
		{
			input:  []int{1, 2},
			output: 1,
		},
		{
			input:  []int{31, 26, 33, 21, 40},
			output: 5,
		},
	}
	testAll(data, t)
}
