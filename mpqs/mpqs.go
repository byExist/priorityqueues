package mpqs

import (
	"cmp"
	"container/heap"
)

// Elem represents an element in the priority queue with an item, priority, and sequence number.
type Elem[T any, P cmp.Ordered] struct {
	item T
	prio P
	seq  int
}

// Item returns the item stored in the element.
func (e Elem[T, P]) Item() T {
	return e.item
}

// Priority returns the priority of the element.
func (e Elem[T, P]) Priority() P {
	return e.prio
}

// Sequence returns the sequence number of the element.
func (e Elem[T, P]) Sequence() int {
	return e.seq
}

type heapImpl[T any, P cmp.Ordered] struct {
	elems    []Elem[T, P]
	lessFunc func(i, j Elem[T, P]) bool
}

func (h *heapImpl[T, P]) Len() int {
	return len(h.elems)
}

func (h *heapImpl[T, P]) Less(i, j int) bool {
	return h.lessFunc(h.elems[i], h.elems[j])
}

func (h *heapImpl[T, P]) Swap(i, j int) {
	h.elems[i], h.elems[j] = h.elems[j], h.elems[i]
}

func (h *heapImpl[T, P]) Push(x any) {
	h.elems = append(h.elems, x.(Elem[T, P]))
}

func (h *heapImpl[T, P]) Pop() any {
	old := h.elems
	n := len(old)
	item := old[n-1]
	h.elems = old[0 : n-1]
	return item
}

func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// PriorityQueue represents a generic priority queue with elements of type T and priority of type P.
type PriorityQueue[T any, P cmp.Ordered] struct {
	heap    *heapImpl[T, P]
	counter func() int
}

// MinFirst compares two elements and returns true if x has lower priority than y.
// Used for min-priority queues.
func MinFirst[T any, P cmp.Ordered](x, y Elem[T, P]) bool {
	return x.prio < y.prio
}

// MaxFirst compares two elements and returns true if x has higher priority than y.
// Used for max-priority queues.
func MaxFirst[T any, P cmp.Ordered](x, y Elem[T, P]) bool {
	return x.prio > y.prio
}

// StableMinFirst compares two elements and returns true if x has lower priority than y,
// or if priorities are equal, if x was inserted earlier (lower sequence number).
func StableMinFirst[T any, P cmp.Ordered](x, y Elem[T, P]) bool {
	if x.prio == y.prio {
		return x.seq < y.seq
	}
	return x.prio < y.prio
}

// StableMaxFirst compares two elements and returns true if x has higher priority than y,
// or if priorities are equal, if x was inserted earlier (lower sequence number).
func StableMaxFirst[T any, P cmp.Ordered](x, y Elem[T, P]) bool {
	if x.prio == y.prio {
		return x.seq < y.seq
	}
	return x.prio > y.prio
}

// New creates a new PriorityQueue with the provided less function.
// The lessFunc determines the priority order: it should return true if x has higher priority than y.
// For a min-priority queue, use: func(x, y Elem[T, P]) bool { return x.prio < y.prio }
func New[T any, P cmp.Ordered](
	lessFunc func(x, y Elem[T, P]) bool,
) *PriorityQueue[T, P] {
	return &PriorityQueue[T, P]{
		heap: &heapImpl[T, P]{
			elems:    []Elem[T, P]{},
			lessFunc: lessFunc,
		},
		counter: counter(),
	}
}

// Clear removes all elements from the priority queue and resets its sequence counter.
func Clear[T any, P cmp.Ordered](pq *PriorityQueue[T, P]) {
	pq.heap.elems = []Elem[T, P]{}
	pq.counter = counter()
}

// Enqueue inserts a new item with the given priority into the priority queue.
func Enqueue[T any, P cmp.Ordered](pq *PriorityQueue[T, P], item T, prio P) {
	elem := Elem[T, P]{
		item: item,
		prio: prio,
		seq:  pq.counter(),
	}
	heap.Push(pq.heap, elem)
}

// Dequeue removes and returns the item with the highest priority from the priority queue.
// The boolean return value indicates whether an item was returned.
func Dequeue[T any, P cmp.Ordered](pq *PriorityQueue[T, P]) (T, bool) {
	if pq.heap.Len() == 0 {
		var zero T
		return zero, false
	}
	elem := heap.Pop(pq.heap).(Elem[T, P])
	return elem.item, true
}

// Peek returns the item with the highest priority without removing it from the queue.
// The boolean return value indicates whether an item was returned.
func Peek[T any, P cmp.Ordered](pq *PriorityQueue[T, P]) (T, bool) {
	if pq.heap.Len() == 0 {
		var zero T
		return zero, false
	}
	return pq.heap.elems[0].item, true
}

// Len returns the number of elements currently in the priority queue.
func Len[T any, P cmp.Ordered](pq *PriorityQueue[T, P]) int {
	return pq.heap.Len()
}
