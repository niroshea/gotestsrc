package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push xxx
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }

//Pop xxx
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// ExampleIntHeap This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func ExampleIntHeap() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		//优先级排序，总是输出最大的或者最小的
		fmt.Printf("%d ", heap.Pop(h))
	}
	// Output:
	// minimum: 1
	// 1 2 3 5
}

func main() {
	ExampleIntHeap()
}
