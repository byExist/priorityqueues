package pqs

import (
	"cmp"
	"container/heap"
)

type heapImpl[T cmp.Ordered] struct {
	items    []T
	lessFunc func(i, j T) bool
}

func (h *heapImpl[T]) Len() int {
	return len(h.items)
}

func (h *heapImpl[T]) Less(i, j int) bool {
	return h.lessFunc(h.items[i], h.items[j])
}

func (h *heapImpl[T]) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *heapImpl[T]) Push(x any) {
	item := x.(T)
	h.items = append(h.items, item)
}

func (h *heapImpl[T]) Pop() any {
	old := h.items
	n := len(old)
	item := old[n-1]
	h.items = old[0 : n-1]
	return item
}

// PriorityQueue represents a generic priority queue data structure.
type PriorityQueue[T cmp.Ordered] struct {
	heap *heapImpl[T]
}

// MinFirst compares two elements and returns true if x has lower priority than y.
// Used for min-priority queues.
func MinFirst[T cmp.Ordered](x, y T) bool {
	return x < y
}

// MaxFirst compares two elements and returns true if x has higher priority than y.
// Used for max-priority queues.
func MaxFirst[T cmp.Ordered](x, y T) bool {
	return x > y
}

// New creates a new PriorityQueue with the provided less function.
// The lessFunc determines the priority order: it should return true if x has higher priority than y.
// For a min-priority queue, use: func(x, y T) bool { return x < y }
func New[T cmp.Ordered](
	lessFunc func(x, y T) bool,
) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		heap: &heapImpl[T]{
			items:    []T{},
			lessFunc: lessFunc,
		},
	}
}

// Clear removes all items from the priority queue.
func Clear[T cmp.Ordered](pq *PriorityQueue[T]) {
	pq.heap.items = []T{}
}

// Enqueue inserts a new item into the priority queue.
func Enqueue[T cmp.Ordered](pq *PriorityQueue[T], item T) {
	heap.Push(pq.heap, item)
}

// Dequeue removes and returns the highest priority item from the priority queue.
// The boolean indicates whether an item was returned.
func Dequeue[T cmp.Ordered](pq *PriorityQueue[T]) (T, bool) {
	if pq.heap.Len() == 0 {
		var zero T
		return zero, false
	}
	elem := heap.Pop(pq.heap).(T)
	return elem, true
}

// Peek returns the highest priority item without removing it from the priority queue.
// The boolean indicates whether an item was returned.
func Peek[T cmp.Ordered](pq *PriorityQueue[T]) (T, bool) {
	if pq.heap.Len() == 0 {
		var zero T
		return zero, false
	}
	elem := pq.heap.items[0]
	return elem, true
}

// Len returns the number of items currently in the priority queue.
func Len[T cmp.Ordered](pq *PriorityQueue[T]) int {
	return pq.heap.Len()
}
