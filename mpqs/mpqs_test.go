package mpqs_test

import (
	"fmt"
	"testing"

	"github.com/byExist/priorityqueues/mpqs"
	"github.com/stretchr/testify/assert"
)

func TestNewAndEnqueueDequeueSingle(t *testing.T) {
	pq := mpqs.New(mpqs.MinFirst[string, int])
	mpqs.Enqueue(pq, "a", 1)
	item, ok := mpqs.Dequeue(pq)
	assert.True(t, ok)
	assert.Equal(t, "a", item)
}

func TestClear(t *testing.T) {
	pq := mpqs.New(mpqs.MinFirst[string, int])
	mpqs.Enqueue(pq, "a", 1)
	mpqs.Clear(pq)
	assert.Equal(t, 0, mpqs.Len(pq))
	_, ok := mpqs.Peek(pq)
	assert.False(t, ok)
}

func TestEnqueueDequeueMultiple(t *testing.T) {
	pq := mpqs.New(mpqs.MinFirst[string, int])
	mpqs.Enqueue(pq, "a", 3)
	mpqs.Enqueue(pq, "b", 1)
	mpqs.Enqueue(pq, "c", 2)

	item1, _ := mpqs.Dequeue(pq)
	item2, _ := mpqs.Dequeue(pq)
	item3, _ := mpqs.Dequeue(pq)

	assert.Equal(t, "b", item1)
	assert.Equal(t, "c", item2)
	assert.Equal(t, "a", item3)
}

func TestEmptyDequeue(t *testing.T) {
	pq := mpqs.New(mpqs.MinFirst[string, int])
	_, ok := mpqs.Dequeue(pq)
	assert.False(t, ok)
}

func TestPeek(t *testing.T) {
	pq := mpqs.New(mpqs.MinFirst[string, int])
	mpqs.Enqueue(pq, "a", 2)
	item, ok := mpqs.Peek(pq)
	assert.True(t, ok)
	assert.Equal(t, "a", item)
}

func TestEmptyPeek(t *testing.T) {
	pq := mpqs.New(mpqs.MinFirst[string, int])
	_, ok := mpqs.Peek(pq)
	assert.False(t, ok)
}

func TestLen(t *testing.T) {
	pq := mpqs.New(mpqs.MinFirst[string, int])
	assert.Equal(t, 0, mpqs.Len(pq))
	mpqs.Enqueue(pq, "a", 1)
	assert.Equal(t, 1, mpqs.Len(pq))
}

func TestMaxFirstOrder(t *testing.T) {
	pq := mpqs.New(mpqs.MaxFirst[string, int])
	mpqs.Enqueue(pq, "low", 1)
	mpqs.Enqueue(pq, "high", 3)
	item, _ := mpqs.Dequeue(pq)
	assert.Equal(t, "high", item)
}

func TestStableMinFirstOrder(t *testing.T) {
	pq := mpqs.New(mpqs.StableMinFirst[string, int])
	mpqs.Enqueue(pq, "a", 1)
	mpqs.Enqueue(pq, "b", 1)
	mpqs.Enqueue(pq, "c", 1)

	first, _ := mpqs.Dequeue(pq)
	second, _ := mpqs.Dequeue(pq)
	third, _ := mpqs.Dequeue(pq)

	assert.Equal(t, "a", first)
	assert.Equal(t, "b", second)
	assert.Equal(t, "c", third)
}

func TestStableMaxFirstOrder(t *testing.T) {
	pq := mpqs.New(mpqs.StableMaxFirst[string, int])
	mpqs.Enqueue(pq, "a", 1)
	mpqs.Enqueue(pq, "b", 1)
	mpqs.Enqueue(pq, "c", 1)

	first, _ := mpqs.Dequeue(pq)
	second, _ := mpqs.Dequeue(pq)
	third, _ := mpqs.Dequeue(pq)

	assert.Equal(t, "a", first)
	assert.Equal(t, "b", second)
	assert.Equal(t, "c", third)
}

func Example_reversedStableMinHeap() {
	reversedStable := func(x, y mpqs.Elem[string, int]) bool {
		if x.Priority() == y.Priority() {
			return x.Sequence() > y.Sequence()
		}
		return x.Priority() < y.Priority()
	}

	pq := mpqs.New(reversedStable)

	mpqs.Enqueue(pq, "a", 1)
	mpqs.Enqueue(pq, "b", 2)
	mpqs.Enqueue(pq, "c", 2)

	first, _ := mpqs.Dequeue(pq)
	second, _ := mpqs.Dequeue(pq)
	third, _ := mpqs.Dequeue(pq)

	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)

	// Output:
	// a
	// c
	// b
}

func ExampleNew() {
	pq := mpqs.New(mpqs.MinFirst[string, int])
	mpqs.Enqueue(pq, "task1", 2)
	mpqs.Enqueue(pq, "task2", 1)
	item, _ := mpqs.Dequeue(pq)
	fmt.Println(item)
	// Output: task2
}

func ExampleClear() {
	pq := mpqs.New(mpqs.MinFirst[string, int])
	mpqs.Enqueue(pq, "task", 1)
	mpqs.Clear(pq)
	fmt.Println(mpqs.Len(pq))
	// Output: 0
}

func ExampleEnqueue() {
	pq := mpqs.New(mpqs.MinFirst[string, int])
	mpqs.Enqueue(pq, "task", 10)
	item, _ := mpqs.Peek(pq)
	fmt.Println(item)
	// Output: task
}

func ExampleDequeue() {
	pq := mpqs.New(mpqs.MinFirst[string, int])
	mpqs.Enqueue(pq, "task1", 2)
	mpqs.Enqueue(pq, "task2", 1)
	item, _ := mpqs.Dequeue(pq)
	fmt.Println(item)
	// Output: task2
}

func ExamplePeek() {
	pq := mpqs.New(mpqs.MinFirst[string, int])
	mpqs.Enqueue(pq, "task", 1)
	item, _ := mpqs.Peek(pq)
	fmt.Println(item)
	// Output: task
}

func ExampleLen() {
	pq := mpqs.New(mpqs.MinFirst[string, int])
	fmt.Println(mpqs.Len(pq))
	mpqs.Enqueue(pq, "a", 1)
	fmt.Println(mpqs.Len(pq))
	// Output:
	// 0
	// 1
}
