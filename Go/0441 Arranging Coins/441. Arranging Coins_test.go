package main

import "testing"

type testForm struct {
	n      int
	output int
}

func test(data []testForm, t *testing.T) {
	for _, e := range data {
		got := arrangeCoins(e.n)
		if got != e.output {
			t.Errorf("at input %d got %d expected %d\n", e.n, got, e.output)
		}
	}
}

func Test(t *testing.T) {
	data := []testForm{
		{
			n:      5,
			output: 2,
		},
		{
			n:      8,
			output: 3,
		},
	}
	test(data, t)
}
