package main

import "testing"

type testInOut struct {
	input int
	ouput int
}

func testIt(data []testInOut, t *testing.T) {
	obj := Constructor()

	for _, e := range data {
		myValue := obj.Ping(e.input)
		if myValue != e.ouput {
			t.Errorf("at input %d got %d expected %d\n", e.input, myValue, e.ouput)
		}
	}
}

func Test1(t *testing.T) {
	data := []testInOut{
		{input: 1, ouput: 1},
		{input: 10, ouput: 2},
		{input: 3001, ouput: 3},
		{input: 3002, ouput: 3},
	}
	testIt(data, t)
}

func Test2(t *testing.T) {
	data := []testInOut{
		{input: 642, ouput: 1},
		{input: 1849, ouput: 2},
		{input: 4921, ouput: 1},
		{input: 5936, ouput: 2},
		{input: 5957, ouput: 3},
	}
	testIt(data, t)
}
