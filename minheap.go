package main

import "fmt"

//MinHeap is an array representation of a Heap
type MinHeap struct {
	array []int
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func (h *MinHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]

}

//Insert willl insert a new element to a heap
func (h *MinHeap) Insert(index int) {
	h.array = append(h.array, index)
	h.minHeapUp(len(h.array) - 1)
}

//minHeap is a heap which
func (h *MinHeap) minHeapUp(index int) {
	for h.array[parent(index)] > h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MinHeap) minHeapDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childCompare := 0

	if l <= lastIndex {
		childCompare = l
	} else if h.array[l] < h.array[r] {
		childCompare = l
	} else {
		childCompare = r
	}
	if h.array[l] < h.array[childCompare] {
		h.swap(index, childCompare)
		index = childCompare

		l, r = left(index), right(index)
	} else {
		return
	}
}

func (h *MinHeap) Extract() int {
	extracted := h.array[0]
	l := len(h.array) - 1
	if len(h.array) == 0 {
		fmt.Println("No Heap Present")
		return -1
	}
	h.array[0] = h.array[l]
	h.array = h.array[:l]
	h.minHeapDown(0)
	return extracted
}

func main() {
	m := &MinHeap{}
	fmt.Println(m)
	buildHeap := []int{10, 20, 11, 9, 8, 4, 2, 5, -1000}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}
	for i := 0; i <= 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
