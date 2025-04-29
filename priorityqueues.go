package priorityqueues

import (
	"cmp"
	"container/heap"
	"iter"
)

type priorityFunc[T any, P cmp.Ordered] func(T) P

// PriorityQueue is a generic priority queue that holds elements of any type T
// and prioritizes them using a user-defined priority function that returns a value of type P.
type PriorityQueue[T any, P cmp.Ordered] struct {
	items        []entry[T, P]
	priorityFunc priorityFunc[T, P]
}

type entry[T any, P cmp.Ordered] struct {
	value    T
	priority P
	index    int
}

type adapter[T any, P cmp.Ordered] struct {
	pq *PriorityQueue[T, P]
}

func (a adapter[T, P]) Len() int {
	return len(a.pq.items)
}

func (a adapter[T, P]) Less(i, j int) bool {
	return a.pq.items[i].priority > a.pq.items[j].priority
}

func (a adapter[T, P]) Swap(i, j int) {
	a.pq.items[i], a.pq.items[j] = a.pq.items[j], a.pq.items[i]
	a.pq.items[i].index = i
	a.pq.items[j].index = j
}

func (a *adapter[T, P]) Push(x any) {
	n := len(a.pq.items)
	item := x.(entry[T, P])
	item.index = n
	a.pq.items = append(a.pq.items, item)
}

func (a *adapter[T, P]) Pop() any {
	old := a.pq.items
	n := len(old)
	item := old[n-1]
	old[n-1] = entry[T, P]{}
	a.pq.items = old[0 : n-1]
	return item
}

// New creates a new empty priority queue with the given priority function.
func New[T any, P cmp.Ordered](priorityFunc priorityFunc[T, P]) *PriorityQueue[T, P] {
	return &PriorityQueue[T, P]{
		items:        make([]entry[T, P], 0),
		priorityFunc: priorityFunc,
	}
}

// FromSeq constructs a priority queue from an iter.Seq and a priority function.
func FromSeq[T any, P cmp.Ordered](seq iter.Seq[T], priorityFunc priorityFunc[T, P]) *PriorityQueue[T, P] {
	pq := New(priorityFunc)
	for v := range seq {
		Enqueue(pq, v)
	}
	return pq
}

// Clone creates a deep copy of the given priority queue.
func Clone[T any, P cmp.Ordered](pq *PriorityQueue[T, P]) *PriorityQueue[T, P] {
	newItems := make([]entry[T, P], len(pq.items))
	copy(newItems, pq.items)
	return &PriorityQueue[T, P]{
		items:        newItems,
		priorityFunc: pq.priorityFunc,
	}
}

// Enqueue inserts a new element into the priority queue.
func Enqueue[T any, P cmp.Ordered](pq *PriorityQueue[T, P], value T) {
	adapter := &adapter[T, P]{pq}
	e := entry[T, P]{
		value:    value,
		priority: pq.priorityFunc(value),
	}
	heap.Push(adapter, e)
}

// Dequeue removes and returns the element with the highest priority.
// Returns false if the queue is empty.
func Dequeue[T any, P cmp.Ordered](pq *PriorityQueue[T, P]) (T, bool) {
	if Len(pq) == 0 {
		var zero T
		return zero, false
	}
	adapter := &adapter[T, P]{pq}
	item := heap.Pop(adapter).(entry[T, P])
	return item.value, true
}

// Peek returns the element with the highest priority without removing it.
// Returns false if the queue is empty.
func Peek[T any, P cmp.Ordered](pq *PriorityQueue[T, P]) (T, bool) {
	if len(pq.items) == 0 {
		var zero T
		return zero, false
	}
	return pq.items[0].value, true
}

// Len returns the number of elements in the priority queue.
func Len[T any, P cmp.Ordered](pq *PriorityQueue[T, P]) int {
	return len(pq.items)
}

// Values returns an iter.Seq of all values in the priority queue,
// without guaranteeing any specific order.
func Values[T any, P cmp.Ordered](pq *PriorityQueue[T, P]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, e := range pq.items {
			if !yield(e.value) {
				break
			}
		}
	}
}

// Clear removes all elements from the priority queue.
func Clear[T any, P cmp.Ordered](pq *PriorityQueue[T, P]) {
	pq.items = pq.items[:0]
}
