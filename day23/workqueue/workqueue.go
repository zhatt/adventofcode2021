// Based on https://pkg.go.dev/container/heap

package workqueue

import (
	"container/heap"
)

// Item is the structure the user manipulates.  The WorkQueue is a heap sorted
// by minimum distance.
type Item struct {
	Item  interface{}
	Score int

	// The index is needed by Update()
	index int // The index of the item in the heap.
}

func NewItem(item interface{}, score int) *Item {
	return &Item{
		Item:  item,
		Score: score,
		index: -1,
	}
}

type WorkQueue struct {
	heap workQueueHeap
}

func NewWorkQueue() WorkQueue {
	wq := WorkQueue{
		heap: make([]*Item, 0),
	}
	return wq
}

func (wq *WorkQueue) Push(work *Item) {
	heap.Push(&wq.heap, work)
}

func (wq *WorkQueue) Pop() *Item {
	return heap.Pop(&wq.heap).(*Item)
}

func (wq *WorkQueue) Len() int {
	return len(wq.heap)
}

func (wq *WorkQueue) Update(item *Item) {
	heap.Fix(&wq.heap, item.index)
}

type workQueueHeap []*Item

func (h workQueueHeap) Len() int { return len(h) }

func (h workQueueHeap) Less(i, j int) bool {
	return h[i].Score < h[j].Score
}

func (h workQueueHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *workQueueHeap) Push(x interface{}) {
	n := len(*h)
	item := x.(*Item)
	item.index = n
	*h = append(*h, item)
}

func (h *workQueueHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*h = old[0 : n-1]
	return item
}
