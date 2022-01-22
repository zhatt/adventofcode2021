// Based on https://pkg.go.dev/container/heap

package priorityqueue

import (
	"container/heap"
)

// Item is the structure the user manipulates.  The WorkQueue is a heap sorted
// by minimum distance.
type Item struct {
	Item     interface{}
	priority int

	// The index is needed by Update()
	index int // The index of the item in the heap.
}

func NewItem(item interface{}, priority int) *Item {
	return &Item{
		Item:     item,
		priority: priority,
		index:    -1,
	}
}

type Queue struct {
	heap pqHeap
}

func New() Queue {
	pq := Queue{
		heap: make([]*Item, 0),
	}
	return pq
}

func (pq *Queue) Push(item *Item) {
	heap.Push(&pq.heap, item)
}

func (pq *Queue) Pop() *Item {
	return heap.Pop(&pq.heap).(*Item)
}

func (pq *Queue) Len() int {
	return len(pq.heap)
}

func (pq *Queue) Update(item *Item, priority int) {
	item.priority = priority
	heap.Fix(&pq.heap, item.index)
}

type pqHeap []*Item

func (h pqHeap) Len() int { return len(h) }

func (h pqHeap) Less(i, j int) bool {
	return h[i].priority < h[j].priority
}

func (h pqHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *pqHeap) Push(x interface{}) {
	n := len(*h)
	item := x.(*Item)
	item.index = n
	*h = append(*h, item)
}

func (h *pqHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*h = old[0 : n-1]
	return item
}
