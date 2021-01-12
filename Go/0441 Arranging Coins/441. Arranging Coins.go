package main

import (
	"math"
)

func arrangeCoins(n int) int {
	return int(-0.5 + math.Sqrt(0.25+2.0*float64(n)))
}
