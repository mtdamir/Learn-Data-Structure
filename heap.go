package main

import "fmt"

// MaxHeap struct has a slice that holds the array
type MaxHeap struct {
	array []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.MaxHeapifyUp(len(h.array) - 1)
}

// Extract returns the largest key, and removes it from the heap
func (h *MaxHeap) Extract() int {
	if len(h.array) == 0 {
		fmt.Println("cannot extract because array length is 0")
		return -1
	}

	extract := h.array[0]
	lastIndex := len(h.array) - 1

	// Move the last element to the root
	h.array[0] = h.array[lastIndex]
	h.array = h.array[:lastIndex]

	h.MaxHeapifyDown(0) // Reorder the heap

	return extract
}

// MaxHeapifyUp will heapify from bottom top
func (h *MaxHeap) MaxHeapifyUp(index int) {
	for index > 0 && h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// MaxHeapifyDown will heapify top to bottom
func (h *MaxHeap) MaxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0
	// loop while index has at last one child
	for l <= lastIndex {
		if l == lastIndex { // when left child is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { // when left child is larger
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}
		// compare array value of current index to larger child and swap if smaller
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

// Get the parent index
func parent(i int) int {
	return (i - 1) / 2
}

// get the left child index
func left(i int) int {
	return 2*i + 1
}

// get the Right child index
func right(i int) int {
	return 2*i + 2
}

// swap keys in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
	m := &MaxHeap{}
	buildHeap := []int{10, 20, 30, 6, 89, 21, 6, 1, 98}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m.array)
	}

	for i := 0; i < 5; i++ {
		extracted := m.Extract()
		fmt.Println("Extracted:", extracted)
		fmt.Println(m.array)
	}
}
