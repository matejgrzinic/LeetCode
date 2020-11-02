package main

func getMaxLen(nums []int) int {
	firstNegative, zeroPosition, negativeCounter, best := -1, -1, true, 0
	for i, num := range nums {
		if num < 0 {
			negativeCounter = !negativeCounter
			if firstNegative == -1 {
				firstNegative = i
			}
		}
		if num == 0 {
			negativeCounter = true
			firstNegative = -1
			zeroPosition = i
		} else {
			if negativeCounter {
				diff := i - zeroPosition
				if diff > best {
					best = diff
				}
			} else {
				diff := i - firstNegative
				if diff > best {
					best = diff
				}
			}
		}
	}
	return best
}
