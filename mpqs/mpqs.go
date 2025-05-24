package mpqs

import (
	"cmp"
	"container/heap"
)

type Elem[T any, P cmp.Ordered] struct {
	item T
	prio P
	seq  int
}

func (e Elem[T, P]) Item() T {
	return e.item
}

func (e Elem[T, P]) Priority() P {
	return e.prio
}

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

type PriorityQueue[T any, P cmp.Ordered] struct {
	heap    *heapImpl[T, P]
	counter func() int
}

func MinFirst[T any, P cmp.Ordered](x, y Elem[T, P]) bool {
	return x.prio < y.prio
}

func MaxFirst[T any, P cmp.Ordered](x, y Elem[T, P]) bool {
	return x.prio > y.prio
}

func StableMinFirst[T any, P cmp.Ordered](x, y Elem[T, P]) bool {
	if x.prio == y.prio {
		return x.seq < y.seq
	}
	return x.prio < y.prio
}

func StableMaxFirst[T any, P cmp.Ordered](x, y Elem[T, P]) bool {
	if x.prio == y.prio {
		return x.seq < y.seq
	}
	return x.prio > y.prio
}

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

func Clear[T any, P cmp.Ordered](pq *PriorityQueue[T, P]) {
	pq.heap.elems = []Elem[T, P]{}
	pq.counter = counter()
}

func Enqueue[T any, P cmp.Ordered](pq *PriorityQueue[T, P], item T, prio P) {
	elem := Elem[T, P]{
		item: item,
		prio: prio,
		seq:  pq.counter(),
	}
	heap.Push(pq.heap, elem)
}

func Dequeue[T any, P cmp.Ordered](pq *PriorityQueue[T, P]) (T, bool) {
	if pq.heap.Len() == 0 {
		var zero T
		return zero, false
	}
	elem := heap.Pop(pq.heap).(Elem[T, P])
	return elem.item, true
}

func Peek[T any, P cmp.Ordered](pq *PriorityQueue[T, P]) (T, bool) {
	if pq.heap.Len() == 0 {
		var zero T
		return zero, false
	}
	return pq.heap.elems[0].item, true
}

func Len[T any, P cmp.Ordered](pq *PriorityQueue[T, P]) int {
	return pq.heap.Len()
}
