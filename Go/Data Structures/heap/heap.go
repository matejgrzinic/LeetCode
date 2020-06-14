package main

import "fmt"

type Heap struct {
	elements []item
}

type item struct {
	cost int
	y    int
	x    int
}

// Push adds new element to heap
func (h *Heap) Push(i item) {
	h.elements = append(h.elements, i)
	h.siftUp(len(h.elements) - 1)
}

func (h *Heap) siftUp(i int) {
	parent := (i - 1) / 2
	if h.elements[i].cost < h.elements[parent].cost {
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
		if h.elements[l].cost < h.elements[i].cost {
			h.swap(i, l)
			h.siftDown(l)
		}
	}
}

func (h *Heap) swap(i int, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

// Pop return smallest priority element
func (h *Heap) Pop() item {
	n := len(h.elements)
	r := h.elements[0]
	h.swap(0, len(h.elements)-1)
	h.elements = h.elements[:n-1]
	h.siftDown(0)
	return r
}

func main() {
	h := Heap{}
	h.Push(item{1, 0, 0})
	h.Push(item{3, 0, 0})
	h.Push(item{2, 0, 0})
	h.Push(item{4, 0, 0})

	for len(h.elements) > 0 {
		i := h.Pop()
		fmt.Println(i.cost)
	}
	fmt.Println(h.elements)
}
