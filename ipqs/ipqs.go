package ipqs

import (
	"cmp"
	"container/heap"
)

// Entry represents an item in the priority queue with its associated priority,
// current index in the heap, and insertion sequence number for stability.
type Entry[T any, P cmp.Ordered] struct {
	item T
	prio P
	loc  int
	seq  int
}

// Item returns the item stored in this entry.
func (e Entry[T, P]) Item() T {
	return e.item
}

// Priority returns the priority associated with this entry.
func (e Entry[T, P]) Priority() P {
	return e.prio
}

// Sequence returns the insertion sequence number of this entry.
func (e Entry[T, P]) Sequence() int {
	return e.seq
}

func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

type entryHeap[I comparable, T any, P cmp.Ordered] struct {
	indexer    func(T) I
	comparator func(x, y Entry[T, P]) bool

	entries []Entry[T, P]
	lookup  map[I]int
	counter func() int
}

func (h *entryHeap[I, T, P]) Len() int {
	return len(h.entries)
}

func (h *entryHeap[I, T, P]) Less(i, j int) bool {
	return h.comparator(h.entries[i], h.entries[j])
}

func (h *entryHeap[I, T, P]) Swap(i, j int) {
	h.entries[i], h.entries[j] = h.entries[j], h.entries[i]
	h.entries[i].loc = i
	h.entries[j].loc = j
	h.lookup[h.indexer(h.entries[i].item)] = i
	h.lookup[h.indexer(h.entries[j].item)] = j
}

func (h *entryHeap[I, T, P]) Push(x any) {
	entry := x.(Entry[T, P])
	entry.loc = len(h.entries)
	entry.seq = h.counter()
	h.entries = append(h.entries, entry)
	h.lookup[h.indexer(entry.item)] = entry.loc
}

func (h *entryHeap[I, T, P]) Pop() any {
	if len(h.entries) == 0 {
		return nil
	}
	item := h.entries[len(h.entries)-1]
	h.entries = h.entries[:len(h.entries)-1]
	delete(h.lookup, h.indexer(item.item))
	return item
}

// IndexedPriorityQueue is a priority queue that supports efficient updates and deletions
// of items identified by a unique key.
type IndexedPriorityQueue[I comparable, T any, P cmp.Ordered] struct {
	heap *entryHeap[I, T, P]
}

// New creates a new IndexedPriorityQueue using the item itself as the unique identifier.
func New[T comparable, P cmp.Ordered](
	comparator func(x, y Entry[T, P]) bool,
) *IndexedPriorityQueue[T, T, P] {
	return NewWithIndexer(
		func(x T) T { return x },
		comparator,
	)
}

// NewWithIndexer creates a new IndexedPriorityQueue with a custom indexer function
// that extracts a unique key from each item.
//
// Note: The indexer function must return a value that uniquely identifies each item.
// If two different items produce the same index key, the most recently added item
// will overwrite the existing one in the internal state. This may lead to unexpected
// behavior during updates or deletions. Ensure the key is unique for each logical item.
func NewWithIndexer[I comparable, T any, P cmp.Ordered](
	indexer func(T) I,
	comparator func(x, y Entry[T, P]) bool,
) *IndexedPriorityQueue[I, T, P] {
	h := &entryHeap[I, T, P]{
		entries:    make([]Entry[T, P], 0),
		lookup:     make(map[I]int),
		counter:    counter(),
		comparator: comparator,
		indexer:    indexer,
	}

	return &IndexedPriorityQueue[I, T, P]{heap: h}
}

// Clear removes all entries from the priority queue.
func Clear[I comparable, T any, P cmp.Ordered](pq *IndexedPriorityQueue[I, T, P]) {
	pq.heap.entries = make([]Entry[T, P], 0)
	pq.heap.lookup = make(map[I]int)
}

// Enqueue adds a new item with the given priority to the priority queue.
// Returns false if the item already exists.
func Enqueue[I comparable, T any, P cmp.Ordered](pq *IndexedPriorityQueue[I, T, P], item T, prio P) bool {
	if Contains(pq, item) {
		return false
	}
	entry := Entry[T, P]{item: item, prio: prio}
	heap.Push(pq.heap, entry)
	return true
}

// Dequeue removes and returns the highest priority item from the queue.
// It returns (zero value, false) if the queue is empty.
func Dequeue[I comparable, T any, P cmp.Ordered](pq *IndexedPriorityQueue[I, T, P]) (T, bool) {
	if pq.heap.Len() == 0 {
		var zero T
		return zero, false
	}
	entry := heap.Pop(pq.heap).(Entry[T, P])
	delete(pq.heap.lookup, pq.heap.indexer(entry.item))
	return entry.item, true
}

// Len returns the number of items currently in the priority queue.
func Len[I comparable, T any, P cmp.Ordered](pq *IndexedPriorityQueue[I, T, P]) int {
	return pq.heap.Len()
}

// Contains checks if the given item is present in the priority queue.
func Contains[I comparable, T any, P cmp.Ordered](pq *IndexedPriorityQueue[I, T, P], item T) bool {
	key := pq.heap.indexer(item)
	_, exists := pq.heap.lookup[key]
	return exists
}

// Update modifies the priority of an existing item in the priority queue.
// Returns false if the item was not found.
func Update[I comparable, T any, P cmp.Ordered](pq *IndexedPriorityQueue[I, T, P], item T, prio P) bool {
	idx := pq.heap.indexer(item)
	loc, exists := pq.heap.lookup[idx]
	if !exists {
		return false
	}
	pq.heap.entries[loc].item = item
	oldprio := pq.heap.entries[loc].prio
	if oldprio != prio {
		pq.heap.entries[loc].prio = prio
		heap.Fix(pq.heap, loc)
	}
	return true
}

// Upsert inserts the item if it does not exist, or updates its priority if it does.
func Upsert[I comparable, T any, P cmp.Ordered](pq *IndexedPriorityQueue[I, T, P], item T, prio P) {
	if Contains(pq, item) {
		Update(pq, item, prio)
	} else {
		Enqueue(pq, item, prio)
	}
}

// Peek returns the highest priority item without removing it from the queue.
// It returns (zero value, false) if the queue is empty.
func Peek[I comparable, T any, P cmp.Ordered](pq *IndexedPriorityQueue[I, T, P]) (T, bool) {
	if pq.heap.Len() == 0 {
		var zero T
		return zero, false
	}
	return pq.heap.entries[0].item, true
}

// Delete removes the specified item from the priority queue by its identifier key.
// Returns false if the item is not found.
func Delete[I comparable, T any, P cmp.Ordered](pq *IndexedPriorityQueue[I, T, P], item T) bool {
	key := pq.heap.indexer(item)
	loc, exists := pq.heap.lookup[key]
	if !exists {
		return false
	}
	heap.Remove(pq.heap, loc)
	delete(pq.heap.lookup, key)
	return true
}
