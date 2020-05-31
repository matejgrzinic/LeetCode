package main

import (
	"./test0064"
)

func main() {
	test0064.Test(minPathSum)
}

func minPathSum(grid [][]int) int {
	h := Heap{}
	n, m := len(grid)-1, len(grid[0])-1
	h.Push(Item{0, 0, 0})

	seen := make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		seen[i] = make([]bool, m+1)
	}
	for {
		item := h.Pop()
		if seen[item.y][item.x] {
			continue
		}
		seen[item.y][item.x] = true
		newCost := item.cost + grid[item.y][item.x]
		if item.y == n && item.x == m {
			return newCost
		}
		if item.y < n && !seen[item.y+1][item.x] {
			h.Push(Item{newCost, item.y + 1, item.x})
		}
		if item.x < m && !seen[item.y][item.x+1] {
			h.Push(Item{newCost, item.y, item.x + 1})
		}
	}
}

type Heap struct {
	elements []Item
}

type Item struct {
	cost int
	y    int
	x    int
}

// Push adds new element to heap
func (h *Heap) Push(i Item) {
	h.elements = append(h.elements, i)
	h.siftUp(len(h.elements) - 1)
}

func (h *Heap) siftUp(i int) {
	parent := (i - 1) / 2
	if h.elements[i].compare(&h.elements[parent]) {
		h.swap(i, parent)
		h.siftUp(parent)
	}
}

func (h *Heap) siftDown(i int) {
	l, r := 2*i+1, 2*i+2
	n := len(h.elements)
	if l < n {
		if r < n && h.elements[r].cost < h.elements[l].cost {
			l = r
		}
		if h.elements[l].compare(&h.elements[i]) {
			h.swap(i, l)
			h.siftDown(l)
		}
	}
}

func (i1 *Item) compare(i2 *Item) bool {
	return i1.cost < i2.cost
}

func (h *Heap) swap(i int, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

// Pop return smallest priority element
func (h *Heap) Pop() Item {
	n := len(h.elements)
	r := h.elements[0]
	h.swap(0, len(h.elements)-1)
	h.elements = h.elements[:n-1]
	h.siftDown(0)
	return r
}
