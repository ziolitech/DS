package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j] // Max Heap
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)

	item := old[n-1]
	*h = old[:n-1]

	return item
}

// using the heap implementation.
func minOperationsHeapBatch(heights []int) int {
	h := MaxHeap(append([]int(nil), heights...))
	heap.Init(&h)

	operations := 0

	for h.Len() > 1 {
		maxHeight := heap.Pop(&h).(int)
		count := 1

		// Extract all piles having the same maximum height.
		for h.Len() > 0 && h[0] == maxHeight {
			heap.Pop(&h)
			count++
		}

		// No lower height exists. All piles are equal.
		if h.Len() == 0 {
			break
		}

		// Next distinct maximum is the target height.
		nextHeight := h[0]

		// All extracted maximum piles become nextHeight.
		for i := 0; i < count; i++ {
			heap.Push(&h, nextHeight)
		}

		operations += count
	}

	return operations
}

func main() {
	fmt.Println("Hello mayank")
	fmt.Println(minOperationsHeapBatch([]int{5, 2, 1}))
	fmt.Println(minOperationsHeapBatch([]int{4, 5, 5, 2, 4}))
}
