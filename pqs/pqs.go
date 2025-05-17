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

type PriorityQueue[T cmp.Ordered] struct {
	heap *heapImpl[T]
}

func MinFirst[T cmp.Ordered](x, y T) bool {
	return x < y
}

func MaxFirst[T cmp.Ordered](x, y T) bool {
	return x > y
}

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

func Clear[T cmp.Ordered](pq *PriorityQueue[T]) {
	pq.heap.items = []T{}
}

func Enqueue[T cmp.Ordered](pq *PriorityQueue[T], item T) {
	heap.Push(pq.heap, item)
}

func Dequeue[T cmp.Ordered](pq *PriorityQueue[T]) (T, bool) {
	if pq.heap.Len() == 0 {
		var zero T
		return zero, false
	}
	elem := heap.Pop(pq.heap).(T)
	return elem, true
}

func Peek[T cmp.Ordered](pq *PriorityQueue[T]) (T, bool) {
	if pq.heap.Len() == 0 {
		var zero T
		return zero, false
	}
	elem := pq.heap.items[0]
	return elem, true
}

func Len[T cmp.Ordered](pq *PriorityQueue[T]) int {
	return pq.heap.Len()
}
