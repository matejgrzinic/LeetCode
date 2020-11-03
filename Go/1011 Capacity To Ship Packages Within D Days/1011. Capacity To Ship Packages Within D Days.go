package main

func shipWithinDays(weights []int, D int) int {
	capacity := 0
	for _, w := range weights {
		if w > capacity {
			capacity = w
		}
	}

	l, r := capacity, capacity
	for !possibleCapacity(weights, D, r) {
		l = r
		r *= 2
	}

	for {
		mid := (l + r + 1) / 2
		if r == mid {
			return mid
		}
		if possibleCapacity(weights, D, mid) {
			r = mid
		} else {
			l = mid
		}
	}
}

func possibleCapacity(weights []int, D int, capacity int) bool {
	loaded := 0
	for _, w := range weights {
		loaded += w
		if loaded > capacity {
			D--
			if D <= 0 {
				return false
			}
			loaded = w
		}
	}
	return true
}
