package kpqs

import (
	"cmp"
	"container/heap"
)

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

type PriorityQueue[K comparable, T any, P cmp.Ordered] struct {
	heap    *heapImpl[K, T, P]
	counter func() int
}

func MinFirst[K comparable, T any, P cmp.Ordered](x, y Elem[T, P]) bool {
	return x.prio < y.prio
}

func MaxFirst[K comparable, T, P cmp.Ordered](x, y Elem[T, P]) bool {
	return x.prio > y.prio
}

func StableMinFirst[K comparable, T, P cmp.Ordered](x, y Elem[T, P]) bool {
	if x.prio == y.prio {
		return x.seq < y.seq
	}
	return x.prio < y.prio
}

func StableMaxFirst[K comparable, T, P cmp.Ordered](x, y Elem[T, P]) bool {
	if x.prio == y.prio {
		return x.seq < y.seq
	}
	return x.prio > y.prio
}

func New[K comparable, T any, P cmp.Ordered](
	keyFunc func(T) K,
	lessFunc func(x, y Elem[T, P]) bool,
) *PriorityQueue[K, T, P] {
	return &PriorityQueue[K, T, P]{
		heap: &heapImpl[K, T, P]{
			elems:    []Elem[T, P]{},
			lookup:   make(map[K]int),
			keyFunc:  keyFunc,
			lessFunc: lessFunc,
		},
		counter: counter(),
	}
}

func Clear[K comparable, T any, P cmp.Ordered](pq *PriorityQueue[K, T, P]) {
	pq.heap.elems = []Elem[T, P]{}
	pq.heap.lookup = make(map[K]int)
	pq.counter = counter()
}

func Enqueue[K comparable, T any, P cmp.Ordered](pq *PriorityQueue[K, T, P], item T, prio P) {
	elem := Elem[T, P]{
		item: item,
		prio: prio,
		seq:  pq.counter(),
	}
	heap.Push(pq.heap, elem)
}

func Dequeue[K comparable, T any, P cmp.Ordered](pq *PriorityQueue[K, T, P]) (T, bool) {
	if pq.heap.Len() == 0 {
		var zero T
		return zero, false
	}
	elem := heap.Pop(pq.heap).(Elem[T, P])
	return elem.item, true
}

func Peek[K comparable, T any, P cmp.Ordered](pq *PriorityQueue[K, T, P]) (T, bool) {
	if pq.heap.Len() == 0 {
		var zero T
		return zero, false
	}
	elem := pq.heap.elems[0]
	return elem.item, true
}

func Move[K comparable, T any, P cmp.Ordered](pq *PriorityQueue[K, T, P], item T, newPrio P) bool {
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

func Delete[K comparable, T any, P cmp.Ordered](pq *PriorityQueue[K, T, P], item T) bool {
	key := pq.heap.keyFunc(item)
	loc, exists := pq.heap.lookup[key]
	if !exists {
		return false
	}
	heap.Remove(pq.heap, loc)
	return true
}

func Len[K comparable, T any, P cmp.Ordered](pq *PriorityQueue[K, T, P]) int {
	return pq.heap.Len()
}

func Contains[K comparable, T any, P cmp.Ordered](pq *PriorityQueue[K, T, P], item T) bool {
	key := pq.heap.keyFunc(item)
	_, exists := pq.heap.lookup[key]
	return exists
}
