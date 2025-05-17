package apqs

import (
	"cmp"
	"container/heap"
)

type Elem[T comparable, P cmp.Ordered] struct {
	item T
	prio P
	seq  int
}

type heapImpl[T comparable, P cmp.Ordered] struct {
	elems  []Elem[T, P]
	lookup map[T]int

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
	h.lookup[h.elems[i].item] = i
	h.lookup[h.elems[j].item] = j
}

func (h *heapImpl[T, P]) Push(x any) {
	item := x.(Elem[T, P])
	h.lookup[item.item] = len(h.elems)
	h.elems = append(h.elems, item)
}

func (h *heapImpl[T, P]) Pop() any {
	old := h.elems
	n := len(old)
	item := old[n-1]
	h.elems = old[0 : n-1]
	delete(h.lookup, item.item)
	return item
}

func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

type PriorityQueue[T comparable, P cmp.Ordered] struct {
	heap     *heapImpl[T, P]
	counter  func() int
	prioFunc func(T) P
}

func MinFirst[T comparable, P cmp.Ordered](x, y Elem[T, P]) bool {
	return x.prio < y.prio
}

func MaxFirst[T comparable, P cmp.Ordered](x, y Elem[T, P]) bool {
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

func New[T comparable, P cmp.Ordered](
	lessFunc func(x, y Elem[T, P]) bool,
	prioFunc func(T) P,
) *PriorityQueue[T, P] {
	return &PriorityQueue[T, P]{
		heap: &heapImpl[T, P]{
			elems:    []Elem[T, P]{},
			lookup:   make(map[T]int),
			lessFunc: lessFunc,
		},
		counter: counter(),
	}
}

func Clear[T comparable, P cmp.Ordered](pq *PriorityQueue[T, P]) {
	pq.heap.elems = []Elem[T, P]{}
	pq.heap.lookup = make(map[T]int)
	pq.counter = counter()
}

func Enqueue[T comparable, P cmp.Ordered](pq *PriorityQueue[T, P], item T) {
	elem := Elem[T, P]{
		item: item,
		prio: pq.prioFunc(item),
		seq:  pq.counter(),
	}
	heap.Push(pq.heap, elem)
}

func Dequeue[T comparable, P cmp.Ordered](pq *PriorityQueue[T, P]) (T, bool) {
	if pq.heap.Len() == 0 {
		var zero T
		return zero, false
	}
	elem := heap.Pop(pq.heap).(Elem[T, P])
	return elem.item, true
}

func Peek[T comparable, P cmp.Ordered](pq *PriorityQueue[T, P]) (T, bool) {
	if pq.heap.Len() == 0 {
		var zero T
		return zero, false
	}
	elem := pq.heap.elems[0]
	return elem.item, true
}

func Refresh[T comparable, P cmp.Ordered](pq *PriorityQueue[T, P], item T) bool {
	loc, exists := pq.heap.lookup[item]
	if !exists {
		return false
	}
	elem := Elem[T, P]{
		item: item,
		prio: pq.prioFunc(item),
		seq:  pq.counter(),
	}
	pq.heap.elems[loc] = elem
	heap.Fix(pq.heap, loc)
	return true
}

func Delete[T comparable, P cmp.Ordered](pq *PriorityQueue[T, P], item T) bool {
	loc, exists := pq.heap.lookup[item]
	if !exists {
		return false
	}
	heap.Remove(pq.heap, loc)
	return true
}

func Len[T comparable, P cmp.Ordered](pq *PriorityQueue[T, P]) int {
	return pq.heap.Len()
}

func Contains[T comparable, P cmp.Ordered](pq *PriorityQueue[T, P], item T) bool {
	_, exists := pq.heap.lookup[item]
	return exists
}
