package main

import "fmt"

//MaxHeap Struct has slice that hols array
type MaxHeap struct {
	array []int
}

//Insert will add an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapUp(len(h.array) - 1)
}

//Extract will extract the node form top up and delete the highest node
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	l := len(h.array) - 1

	if len(h.array) == 0 {
		fmt.Println("The Heap is Empty ")
		return -1
	}
	h.array[0] = h.array[l]
	h.array = h.array[:l]

	h.maxHeapDown(0)

	return extracted
}

//MaxHeap will return highest heap from bottom top
func (h *MaxHeap) maxHeapUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) maxHeapDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childCompare := 0

	for l <= lastIndex {
		if l == lastIndex {
			childCompare = l
		} else if h.array[l] > h.array[r] {
			childCompare = l
		} else {
			childCompare = r
		}
		if h.array[index] < h.array[childCompare] {
			h.swap(index, childCompare)
			index = childCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}

}

//Parent Index
func parent(i int) int {
	return (i - 1) / 2
}

//left child index
func left(i int) int {
	return 2*i + 1
}

//right child index
func right(i int) int {
	return 2*i + 2
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]

}

func main() {
	m := MaxHeap{}
	buildHeap := []int{10, 20, 30, 50, 55, 5, 60, 43, 65, 67, 99, -1000}

	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}
	for i := 0; i <= 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
