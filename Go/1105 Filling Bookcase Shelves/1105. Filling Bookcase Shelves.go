package main

import (
	"math"
)

func minHeightShelves(books [][]int, shelfWidth int) int {
	dp := make([]int, len(books)+1)

	for i := 1; i < len(dp); i++ {
		maxHeight, widthUsed := 0, 0
		dp[i] = math.MaxInt32

		for j := i - 1; j >= 0; j-- {
			bookWidth, bookHeight := books[j][0], books[j][1]

			widthUsed += bookWidth
			if bookHeight > maxHeight {
				maxHeight = bookHeight
			}

			if widthUsed <= shelfWidth {
				x := dp[j] + maxHeight
				if x < dp[i] {
					dp[i] = x
				}
			} else {
				break
			}
		}
	}
	return dp[len(dp)-1]
}
