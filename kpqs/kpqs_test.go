package kpqs_test

import (
	"fmt"
	"testing"

	"github.com/byExist/priorityqueues/kpqs"
	"github.com/stretchr/testify/assert"
)

type Task struct {
	ID       string
	Priority int
}

func TestClear(t *testing.T) {
	tasks := map[string]*Task{
		"abc": {ID: "abc", Priority: 3},
	}
	pq := kpqs.New(
		kpqs.MinFirst[*Task, int],
		func(t *Task) string { return t.ID },
		func(t *Task) int { return t.Priority },
	)
	kpqs.Enqueue(pq, tasks["abc"])
	kpqs.Clear(pq)
	assert.Equal(t, 0, kpqs.Len(pq))
	_, ok := kpqs.Peek(pq)
	assert.False(t, ok)
}

func TestEnqueueDequeueSingle(t *testing.T) {
	tasks := map[string]*Task{
		"abc": {ID: "abc", Priority: 3},
	}
	pq := kpqs.New(
		kpqs.MinFirst[*Task, int],
		func(t *Task) string { return t.ID },
		func(t *Task) int { return t.Priority },
	)
	kpqs.Enqueue(pq, tasks["abc"])
	item, ok := kpqs.Dequeue(pq)
	assert.True(t, ok)
	assert.Equal(t, "abc", item.ID)
}

func TestEnqueueDequeueMultiple(t *testing.T) {
	tasks := map[string]*Task{
		"longer": {ID: "longer", Priority: 6},
		"a":      {ID: "a", Priority: 1},
		"mid":    {ID: "mid", Priority: 3},
	}
	pq := kpqs.New(
		kpqs.MinFirst[*Task, int],
		func(t *Task) string { return t.ID },
		func(t *Task) int { return t.Priority },
	)
	kpqs.Enqueue(pq, tasks["longer"])
	kpqs.Enqueue(pq, tasks["a"])
	kpqs.Enqueue(pq, tasks["mid"])

	first, _ := kpqs.Dequeue(pq)
	second, _ := kpqs.Dequeue(pq)
	third, _ := kpqs.Dequeue(pq)

	assert.Equal(t, "a", first.ID)
	assert.Equal(t, "mid", second.ID)
	assert.Equal(t, "longer", third.ID)
}

func TestPeek(t *testing.T) {
	tasks := map[string]*Task{
		"abc": {ID: "abc", Priority: 3},
	}
	pq := kpqs.New(
		kpqs.MinFirst[*Task, int],
		func(t *Task) string { return t.ID },
		func(t *Task) int { return t.Priority },
	)
	kpqs.Enqueue(pq, tasks["abc"])
	item, ok := kpqs.Peek(pq)
	assert.True(t, ok)
	assert.Equal(t, "abc", item.ID)
}

func TestUpdate(t *testing.T) {
	t1, t2 := &Task{ID: "task1", Priority: 5}, &Task{ID: "task2", Priority: 6}
	pq := kpqs.New(
		kpqs.MinFirst[*Task, int],
		func(t *Task) string { return t.ID },
		func(t *Task) int { return t.Priority },
	)
	kpqs.Enqueue(pq, t1)
	kpqs.Enqueue(pq, t2)
	t2.Priority = 1
	kpqs.Update(pq, t2)
	item, ok := kpqs.Peek(pq)
	assert.True(t, ok)
	assert.Equal(t, "task2", item.ID)
}

func TestDelete(t *testing.T) {
	task := &Task{ID: "t1", Priority: 1}
	pq := kpqs.New(
		kpqs.MinFirst[*Task, int],
		func(t *Task) string { return t.ID },
		func(t *Task) int { return t.Priority },
	)
	kpqs.Enqueue(pq, task)
	ok := kpqs.Delete(pq, task)
	assert.True(t, ok)
	assert.Equal(t, 0, kpqs.Len(pq))
}

func TestLen(t *testing.T) {
	tasks := map[string]*Task{
		"abc": {ID: "abc", Priority: 3},
	}
	pq := kpqs.New(
		kpqs.MinFirst[*Task, int],
		func(t *Task) string { return t.ID },
		func(t *Task) int { return t.Priority },
	)
	assert.Equal(t, 0, kpqs.Len(pq))
	kpqs.Enqueue(pq, tasks["abc"])
	assert.Equal(t, 1, kpqs.Len(pq))
}

func TestContains(t *testing.T) {
	task := &Task{ID: "t1", Priority: 1}
	pq := kpqs.New(
		kpqs.MinFirst[*Task, int],
		func(t *Task) string { return t.ID },
		func(t *Task) int { return t.Priority },
	)
	assert.False(t, kpqs.Contains(pq, task))
	kpqs.Enqueue(pq, task)
	assert.True(t, kpqs.Contains(pq, task))
}

func TestEmptyDequeuePeek(t *testing.T) {
	pq := kpqs.New(
		kpqs.MinFirst[*Task, int],
		func(t *Task) string { return t.ID },
		func(t *Task) int { return t.Priority },
	)
	_, ok1 := kpqs.Dequeue(pq)
	_, ok2 := kpqs.Peek(pq)
	assert.False(t, ok1)
	assert.False(t, ok2)
}

func ExampleNew() {
	type Task struct {
		ID       string
		Priority int
	}
	keyFunc := func(t *Task) string { return t.ID }
	prioFunc := func(t *Task) int { return t.Priority }

	pq := kpqs.New(
		kpqs.StableMinFirst[*Task, int],
		keyFunc,
		prioFunc,
	)

	kpqs.Enqueue(pq, &Task{ID: "t1", Priority: 2})
	kpqs.Enqueue(pq, &Task{ID: "t2", Priority: 2})
	kpqs.Enqueue(pq, &Task{ID: "t3", Priority: 1})

	first, _ := kpqs.Dequeue(pq)
	second, _ := kpqs.Dequeue(pq)

	fmt.Println(first.ID)
	fmt.Println(second.ID)

	// Output:
	// t3
	// t1
}

func ExampleClear() {
	type Task struct {
		ID       string
		Priority int
	}
	pq := kpqs.New(
		kpqs.MinFirst[*Task, int],
		func(t *Task) string { return t.ID },
		func(t *Task) int { return t.Priority },
	)

	kpqs.Enqueue(pq, &Task{ID: "t1", Priority: 1})
	kpqs.Enqueue(pq, &Task{ID: "t2", Priority: 2})
	fmt.Println("len before:", kpqs.Len(pq))

	kpqs.Clear(pq)
	fmt.Println("len after:", kpqs.Len(pq))

	// Output:
	// len before: 2
	// len after: 0
}

func ExampleEnqueue_dequeuePeek() {
	type Task struct {
		ID       string
		Priority int
	}
	pq := kpqs.New(
		kpqs.MinFirst[*Task, int],
		func(t *Task) string { return t.ID },
		func(t *Task) int { return t.Priority },
	)
	kpqs.Enqueue(pq, &Task{ID: "t1", Priority: 2})
	kpqs.Enqueue(pq, &Task{ID: "t2", Priority: 1})

	item, _ := kpqs.Peek(pq)
	fmt.Println("peek:", item.ID)

	item, _ = kpqs.Dequeue(pq)
	fmt.Println("dequeue:", item.ID)

	// Output:
	// peek: t2
	// dequeue: t2
}

func ExampleUpdate() {
	type Task struct {
		ID       string
		Priority int
	}
	pq := kpqs.New(
		kpqs.MinFirst[*Task, int],
		func(t *Task) string { return t.ID },
		func(t *Task) int { return t.Priority },
	)
	t := &Task{ID: "t1", Priority: 5}
	kpqs.Enqueue(pq, t)
	t.Priority = 1
	kpqs.Update(pq, t)
	item, _ := kpqs.Peek(pq)
	fmt.Println(item.ID)

	// Output:
	// t1
}

func ExampleDelete() {
	type Task struct {
		ID       string
		Priority int
	}
	pq := kpqs.New(
		kpqs.MinFirst[*Task, int],
		func(t *Task) string { return t.ID },
		func(t *Task) int { return t.Priority },
	)
	t := &Task{ID: "t1", Priority: 1}
	kpqs.Enqueue(pq, t)
	ok := kpqs.Delete(pq, t)
	fmt.Println("deleted:", ok)
	fmt.Println("len:", kpqs.Len(pq))

	// Output:
	// deleted: true
	// len: 0
}

func ExampleContains() {
	type Task struct {
		ID       string
		Priority int
	}
	pq := kpqs.New(
		kpqs.MinFirst[*Task, int],
		func(t *Task) string { return t.ID },
		func(t *Task) int { return t.Priority },
	)
	t := &Task{ID: "t1", Priority: 1}
	kpqs.Enqueue(pq, t)
	fmt.Println(kpqs.Contains(pq, t))

	// Output:
	// true
}
