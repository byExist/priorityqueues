package ipqs_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/byExist/priorityqueues/ipqs"

	"github.com/stretchr/testify/require"
)

func TestEnqueueBasic(t *testing.T) {
	pq := ipqs.New(func(x, y ipqs.Entry[int, int]) bool {
		return x.Priority() < y.Priority()
	})

	require.True(t, ipqs.Enqueue(pq, 10, 1))
	require.Equal(t, 1, ipqs.Len(pq))
}

func TestEnqueueDuplicate(t *testing.T) {
	pq := ipqs.New(func(x, y ipqs.Entry[int, int]) bool {
		return x.Priority() < y.Priority()
	})

	require.True(t, ipqs.Enqueue(pq, 10, 1))
	require.False(t, ipqs.Enqueue(pq, 10, 2))

	require.Equal(t, 1, ipqs.Len(pq))

	val, ok := ipqs.Dequeue(pq)
	require.True(t, ok)
	require.Equal(t, 10, val)
}

func TestDequeueEmpty(t *testing.T) {
	pq := ipqs.New(func(x, y ipqs.Entry[int, int]) bool {
		return x.Priority() < y.Priority()
	})

	val, ok := ipqs.Dequeue(pq)
	require.False(t, ok)
	require.Equal(t, 0, val)
}

func TestDequeueOrdering(t *testing.T) {
	pq := ipqs.New(func(x, y ipqs.Entry[int, int]) bool {
		return x.Priority() < y.Priority()
	})

	require.True(t, ipqs.Enqueue(pq, 10, 3))
	require.True(t, ipqs.Enqueue(pq, 20, 1))
	require.True(t, ipqs.Enqueue(pq, 30, 2))

	v1, ok1 := ipqs.Dequeue(pq)
	v2, ok2 := ipqs.Dequeue(pq)
	v3, ok3 := ipqs.Dequeue(pq)

	require.True(t, ok1)
	require.True(t, ok2)
	require.True(t, ok3)

	require.Equal(t, 20, v1)
	require.Equal(t, 30, v2)
	require.Equal(t, 10, v3)
}

func TestPeekMatchesDequeue(t *testing.T) {
	pq := ipqs.New(func(x, y ipqs.Entry[int, int]) bool {
		return x.Priority() < y.Priority()
	})

	require.True(t, ipqs.Enqueue(pq, 42, 100))

	peek, ok1 := ipqs.Peek(pq)
	require.True(t, ok1)

	val, ok2 := ipqs.Dequeue(pq)
	require.True(t, ok2)

	require.Equal(t, peek, val)
}

func TestDeleteExistingItem(t *testing.T) {
	pq := ipqs.New(func(x, y ipqs.Entry[int, int]) bool {
		return x.Priority() < y.Priority()
	})

	require.True(t, ipqs.Enqueue(pq, 5, 2))

	ok := ipqs.Delete(pq, 5)
	require.True(t, ok)
	require.False(t, ipqs.Contains(pq, 5))
}

func TestDeleteNonExistingItem(t *testing.T) {
	pq := ipqs.New(func(x, y ipqs.Entry[int, int]) bool {
		return x.Priority() < y.Priority()
	})

	require.False(t, ipqs.Delete(pq, 999))
}

func TestUpdatePriority(t *testing.T) {
	pq := ipqs.New(func(x, y ipqs.Entry[int, int]) bool {
		return x.Priority() < y.Priority()
	})

	require.True(t, ipqs.Enqueue(pq, 10, 3))
	require.True(t, ipqs.Update(pq, 10, 1))

	val, ok := ipqs.Peek(pq)
	require.True(t, ok)
	require.Equal(t, 10, val)
}

func TestUpdateNonExistingItem(t *testing.T) {
	pq := ipqs.New(func(x, y ipqs.Entry[int, int]) bool {
		return x.Priority() < y.Priority()
	})

	require.False(t, ipqs.Update(pq, 99, 1))
}

func TestUpsertInsertNew(t *testing.T) {
	pq := ipqs.New(func(x, y ipqs.Entry[int, int]) bool {
		return x.Priority() < y.Priority()
	})

	ipqs.Upsert(pq, 1, 5)
	val, ok := ipqs.Dequeue(pq)

	require.True(t, ok)
	require.Equal(t, 1, val)
}

func TestUpsertUpdateExisting(t *testing.T) {
	pq := ipqs.New(func(x, y ipqs.Entry[int, int]) bool {
		return x.Priority() < y.Priority()
	})

	require.True(t, ipqs.Enqueue(pq, 1, 10))
	ipqs.Upsert(pq, 1, 1)

	val, ok := ipqs.Peek(pq)
	require.True(t, ok)
	require.Equal(t, 1, val)
}

func TestClearQueue(t *testing.T) {
	pq := ipqs.New(func(x, y ipqs.Entry[int, int]) bool {
		return x.Priority() < y.Priority()
	})

	require.True(t, ipqs.Enqueue(pq, 1, 1))
	require.True(t, ipqs.Enqueue(pq, 2, 2))
	ipqs.Clear(pq)

	require.Equal(t, 0, ipqs.Len(pq))
	_, ok := ipqs.Dequeue(pq)
	require.False(t, ok)
}

func TestContainsAccuracy(t *testing.T) {
	pq := ipqs.New(func(x, y ipqs.Entry[int, int]) bool {
		return x.Priority() < y.Priority()
	})

	require.False(t, ipqs.Contains(pq, 1))

	require.True(t, ipqs.Enqueue(pq, 1, 1))
	require.True(t, ipqs.Contains(pq, 1))

	ipqs.Delete(pq, 1)
	require.False(t, ipqs.Contains(pq, 1))
}

func TestFIFOWithSamePriority(t *testing.T) {
	pq := ipqs.New(func(x, y ipqs.Entry[int, int]) bool {
		if x.Priority() == y.Priority() {
			return x.Sequence() < y.Sequence()
		}
		return x.Priority() < y.Priority()
	})

	require.True(t, ipqs.Enqueue(pq, 1, 10))
	require.True(t, ipqs.Enqueue(pq, 2, 10))
	require.True(t, ipqs.Enqueue(pq, 3, 10))

	v1, ok1 := ipqs.Dequeue(pq)
	v2, ok2 := ipqs.Dequeue(pq)
	v3, ok3 := ipqs.Dequeue(pq)

	require.True(t, ok1)
	require.True(t, ok2)
	require.True(t, ok3)

	require.Equal(t, 1, v1)
	require.Equal(t, 2, v2)
	require.Equal(t, 3, v3)
}

func TestGenericType_Int(t *testing.T) {
	pq := ipqs.New(func(x, y ipqs.Entry[int, int]) bool {
		return x.Priority() < y.Priority()
	})

	require.True(t, ipqs.Enqueue(pq, 5, 50))
	require.True(t, ipqs.Enqueue(pq, 3, 30))
	require.True(t, ipqs.Enqueue(pq, 4, 40))

	v1, _ := ipqs.Dequeue(pq)
	v2, _ := ipqs.Dequeue(pq)
	v3, _ := ipqs.Dequeue(pq)

	require.Equal(t, 3, v1)
	require.Equal(t, 4, v2)
	require.Equal(t, 5, v3)
}

func TestGenericType_Struct(t *testing.T) {
	type job struct {
		id string
	}
	pq := ipqs.NewWithIndexer(
		func(j job) string { return j.id },
		func(a, b ipqs.Entry[job, int]) bool {
			return a.Priority() < b.Priority()
		},
	)

	require.True(t, ipqs.Enqueue(pq, job{"a"}, 2))
	require.True(t, ipqs.Enqueue(pq, job{"b"}, 1))
	require.True(t, ipqs.Enqueue(pq, job{"c"}, 3))

	j1, _ := ipqs.Dequeue(pq)
	j2, _ := ipqs.Dequeue(pq)
	j3, _ := ipqs.Dequeue(pq)

	require.Equal(t, "b", j1.id)
	require.Equal(t, "a", j2.id)
	require.Equal(t, "c", j3.id)
}

func TestCustomComparator_MaxHeap(t *testing.T) {
	pq := ipqs.New(func(x, y ipqs.Entry[int, int]) bool {
		return x.Priority() > y.Priority() // Max-heap
	})

	require.True(t, ipqs.Enqueue(pq, 1, 10))
	require.True(t, ipqs.Enqueue(pq, 2, 30))
	require.True(t, ipqs.Enqueue(pq, 3, 20))

	v1, ok1 := ipqs.Dequeue(pq)
	v2, ok2 := ipqs.Dequeue(pq)
	v3, ok3 := ipqs.Dequeue(pq)

	require.True(t, ok1)
	require.True(t, ok2)
	require.True(t, ok3)

	require.Equal(t, 2, v1) // highest priority
	require.Equal(t, 3, v2)
	require.Equal(t, 1, v3)
}

func Example_scheduling() {
	type Event struct {
		ID   string
		Time time.Time
	}

	events := []*Event{
		{"A", time.Date(2023, 10, 1, 10, 0, 0, 0, time.UTC)},
		{"B", time.Date(2023, 10, 1, 9, 0, 0, 0, time.UTC)},
		{"C", time.Date(2023, 10, 1, 8, 0, 0, 0, time.UTC)},
		{"D", time.Date(2023, 10, 1, 11, 0, 0, 0, time.UTC)},
	}

	identifier := func(e *Event) string { return e.ID }
	stableMin := func(x, y ipqs.Entry[*Event, int]) bool {
		if x.Priority() == y.Priority() {
			return x.Sequence() < y.Sequence()
		}
		return x.Priority() < y.Priority()
	}
	pq := ipqs.NewWithIndexer(identifier, stableMin)

	for _, e := range events {
		_ = ipqs.Enqueue(pq, e, int(e.Time.Unix()))
	}

	b := events[1]
	b.Time = time.Date(2023, 10, 1, 11, 0, 0, 0, time.UTC)
	ipqs.Update(pq, b, int(b.Time.Unix()))

	for ipqs.Len(pq) > 0 {
		e, _ := ipqs.Dequeue(pq)
		fmt.Print(e.ID, " ")
	}
	// Output: C A B D
}

func Example_dijkstra() {
	type Node string
	type Edge struct {
		To     Node
		Weight int
	}
	graph := map[Node][]Edge{
		"A": {{"B", 1}, {"C", 4}},
		"B": {{"C", 2}, {"D", 5}},
		"C": {{"D", 1}},
		"D": {},
	}

	dist := map[Node]int{"A": 0}
	pq := ipqs.NewWithIndexer(
		func(node Node) Node { return node },
		func(a, b ipqs.Entry[Node, int]) bool { return a.Priority() < b.Priority() },
	)

	ipqs.Enqueue(pq, "A", 0)

	for ipqs.Len(pq) > 0 {
		u, _ := ipqs.Dequeue(pq)
		for _, e := range graph[u] {
			newDist := dist[u] + e.Weight
			if d, ok := dist[e.To]; !ok || newDist < d {
				dist[e.To] = newDist
				ipqs.Upsert(pq, e.To, newDist)
			}
		}
	}

	keys := []Node{"A", "B", "C", "D"}
	for _, k := range keys {
		fmt.Printf("%s:%d ", k, dist[k])
	}
	// Output: A:0 B:1 C:3 D:4
}

func ExampleNew() {
	pq := ipqs.New(func(x, y ipqs.Entry[int, int]) bool {
		return x.Priority() < y.Priority()
	})

	_ = ipqs.Enqueue(pq, 1, 5)
	_ = ipqs.Enqueue(pq, 1, 2)
	_ = ipqs.Enqueue(pq, 2, 3)
	val, _ := ipqs.Dequeue(pq)
	fmt.Println(val)
	// Output: 2
}

func ExampleNewWithIndexer() {
	type Task struct {
		ID   string
		Name string
	}

	pq := ipqs.NewWithIndexer(
		func(t Task) string { return t.ID },
		func(x, y ipqs.Entry[Task, int]) bool {
			return x.Priority() < y.Priority()
		},
	)

	ipqs.Enqueue(pq, Task{"a", "task1"}, 2)
	ipqs.Enqueue(pq, Task{"b", "task2"}, 1)

	// Even if names differ, same ID would be rejected
	replaced := ipqs.Enqueue(pq, Task{"a", "task1-new"}, 0)
	fmt.Println(replaced)

	t, _ := ipqs.Dequeue(pq)
	fmt.Println(t.Name)
	// Output:
	// false
	// task2
}

func ExampleClear() {
	pq := ipqs.New(func(x, y ipqs.Entry[int, int]) bool {
		return x.Priority() < y.Priority()
	})

	_ = ipqs.Enqueue(pq, 1, 1)
	ipqs.Clear(pq)
	fmt.Println(ipqs.Len(pq))
	// Output: 0
}

func ExampleContains() {
	pq := ipqs.New(func(x, y ipqs.Entry[int, int]) bool {
		return x.Priority() < y.Priority()
	})

	_ = ipqs.Enqueue(pq, 1, 1)
	fmt.Println(ipqs.Contains(pq, 1))
	// Output: true
}

func ExampleDelete() {
	pq := ipqs.New(func(x, y ipqs.Entry[int, int]) bool {
		return x.Priority() < y.Priority()
	})

	ipqs.Enqueue(pq, 1, 1)
	fmt.Println(ipqs.Len(pq))
	ipqs.Delete(pq, 1)
	fmt.Println(ipqs.Contains(pq, 1))
	fmt.Println(ipqs.Len(pq))
	// Output:
	// 1
	// false
	// 0
}

func ExampleDequeue() {
	pq := ipqs.New(func(x, y ipqs.Entry[int, int]) bool {
		return x.Priority() < y.Priority()
	})

	_ = ipqs.Enqueue(pq, 1, 3)
	_ = ipqs.Enqueue(pq, 2, 1)
	val, _ := ipqs.Dequeue(pq)
	fmt.Println(val)
	// Output: 2
}

func ExampleEnqueue() {
	pq := ipqs.New(func(x, y ipqs.Entry[int, int]) bool {
		return x.Priority() < y.Priority()
	})

	result := ipqs.Enqueue(pq, 1, 5)
	fmt.Println(result)
	// Output: true
}

func ExampleLen() {
	pq := ipqs.New(func(x, y ipqs.Entry[int, int]) bool {
		return x.Priority() < y.Priority()
	})

	_ = ipqs.Enqueue(pq, 1, 1)
	_ = ipqs.Enqueue(pq, 2, 2)
	fmt.Println(ipqs.Len(pq))
	// Output: 2
}

func ExamplePeek() {
	pq := ipqs.New(func(x, y ipqs.Entry[int, int]) bool {
		return x.Priority() < y.Priority()
	})

	_ = ipqs.Enqueue(pq, 1, 2)
	_ = ipqs.Enqueue(pq, 2, 1)
	val, _ := ipqs.Peek(pq)
	fmt.Println(val)
	// Output: 2
}

func ExampleUpdate() {
	pq := ipqs.New(func(x, y ipqs.Entry[int, int]) bool {
		return x.Priority() < y.Priority()
	})

	ipqs.Enqueue(pq, 1, 5)
	ipqs.Enqueue(pq, 2, 1)
	ipqs.Update(pq, 1, 0) // Now 1 has higher priority
	val, _ := ipqs.Peek(pq)
	fmt.Println(val)
	// Output: 1
}

func ExampleUpsert() {
	pq := ipqs.New(func(x, y ipqs.Entry[int, int]) bool {
		return x.Priority() < y.Priority()
	})

	ipqs.Upsert(pq, 1, 3) // Insert new
	ipqs.Upsert(pq, 2, 2) // Insert another
	ipqs.Upsert(pq, 1, 1) // Update existing

	val, _ := ipqs.Peek(pq)
	fmt.Println(val)
	// Output: 1
}
