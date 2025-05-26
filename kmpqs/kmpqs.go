package kmpqs

import (
	"cmp"
	"container/heap"
)

// Elem represents an element in the priority queue with an item, its priority, and a sequence number.
type Elem[T any, P cmp.Ordered] struct {
	item T
	prio P
	seq  int
}

type heapImpl[K comparable, T any, P cmp.Ordered] struct {
	elems  []Elem[T, P]
	lookup map[K]int

	keyFunc  func(T) K
	lessFunc func(i, j Elem[T, P]) bool
}

func (h *heapImpl[K, T, P]) Len() int {
	return len(h.elems)
}

func (h *heapImpl[K, T, P]) Less(i, j int) bool {
	return h.lessFunc(h.elems[i], h.elems[j])
}

func (h *heapImpl[K, T, P]) Swap(i, j int) {
	h.elems[i], h.elems[j] = h.elems[j], h.elems[i]
	h.lookup[h.keyFunc(h.elems[i].item)] = i
	h.lookup[h.keyFunc(h.elems[j].item)] = j
}

func (h *heapImpl[K, T, P]) Push(x any) {
	item := x.(Elem[T, P])
	h.lookup[h.keyFunc(item.item)] = len(h.elems)
	h.elems = append(h.elems, item)
}

func (h *heapImpl[K, T, P]) Pop() any {
	old := h.elems
	n := len(old)
	item := old[n-1]
	h.elems = old[0 : n-1]
	delete(h.lookup, h.keyFunc(item.item))
	return item
}

func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// PriorityQueue implements a priority queue with efficient update, delete, and lookup operations.
type PriorityQueue[K comparable, T any, P cmp.Ordered] struct {
	heap    *heapImpl[K, T, P]
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
func New[K comparable, T any, P cmp.Ordered](
	lessFunc func(x, y Elem[T, P]) bool,
	keyFunc func(T) K,
) *PriorityQueue[K, T, P] {
	return &PriorityQueue[K, T, P]{
		heap: &heapImpl[K, T, P]{
			elems:    []Elem[T, P]{},
			lookup:   make(map[K]int),
			lessFunc: lessFunc,
			keyFunc:  keyFunc,
		},
		counter: counter(),
	}
}

// Clear removes all elements from the priority queue.
func Clear[K comparable, T any, P cmp.Ordered](pq *PriorityQueue[K, T, P]) {
	pq.heap.elems = []Elem[T, P]{}
	pq.heap.lookup = make(map[K]int)
	pq.counter = counter()
}

// Enqueue inserts a new item with the given priority into the priority queue.
func Enqueue[K comparable, T any, P cmp.Ordered](pq *PriorityQueue[K, T, P], item T, prio P) {
	elem := Elem[T, P]{
		item: item,
		prio: prio,
		seq:  pq.counter(),
	}
	heap.Push(pq.heap, elem)
}

// Dequeue removes and returns the highest priority item from the priority queue.
func Dequeue[K comparable, T any, P cmp.Ordered](pq *PriorityQueue[K, T, P]) (T, bool) {
	if pq.heap.Len() == 0 {
		var zero T
		return zero, false
	}
	elem := heap.Pop(pq.heap).(Elem[T, P])
	return elem.item, true
}

// Peek returns the highest priority item without removing it from the priority queue.
func Peek[K comparable, T any, P cmp.Ordered](pq *PriorityQueue[K, T, P]) (T, bool) {
	if pq.heap.Len() == 0 {
		var zero T
		return zero, false
	}
	elem := pq.heap.elems[0]
	return elem.item, true
}

// Update modifies the priority of an existing item using the queue's prioFunc.
// Returns true if the item exists and was successfully updated.
func Update[K comparable, T any, P cmp.Ordered](pq *PriorityQueue[K, T, P], item T, newPrio P) bool {
	key := pq.heap.keyFunc(item)
	loc, exists := pq.heap.lookup[key]
	if !exists {
		return false
	}
	elem := Elem[T, P]{
		item: item,
		prio: newPrio,
		seq:  pq.counter(),
	}
	pq.heap.elems[loc] = elem
	heap.Fix(pq.heap, loc)
	return true
}

// Delete removes an item identified by its key from the priority queue.
// Returns true if the item existed and was successfully removed.
func Delete[K comparable, T any, P cmp.Ordered](pq *PriorityQueue[K, T, P], item T) bool {
	key := pq.heap.keyFunc(item)
	loc, exists := pq.heap.lookup[key]
	if !exists {
		return false
	}
	heap.Remove(pq.heap, loc)
	return true
}

// Len returns the number of items currently in the priority queue.
func Len[K comparable, T any, P cmp.Ordered](pq *PriorityQueue[K, T, P]) int {
	return pq.heap.Len()
}

// Contains returns true if the queue contains an item identified by its key.
func Contains[K comparable, T any, P cmp.Ordered](pq *PriorityQueue[K, T, P], item T) bool {
	key := pq.heap.keyFunc(item)
	_, exists := pq.heap.lookup[key]
	return exists
}
