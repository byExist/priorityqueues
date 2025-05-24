package pqs_test

import (
	"fmt"
	"testing"

	"github.com/byExist/priorityqueues/pqs"
	"github.com/stretchr/testify/assert"
)

func TestEnqueueDequeueSingle(t *testing.T) {
	pq := pqs.New(pqs.MinFirst[int])
	pqs.Enqueue(pq, 42)
	item, ok := pqs.Dequeue(pq)
	assert.True(t, ok)
	assert.Equal(t, 42, item)
}

func TestEnqueueDequeueMultiple(t *testing.T) {
	pq := pqs.New(pqs.MinFirst[int])
	pqs.Enqueue(pq, 3)
	pqs.Enqueue(pq, 1)
	pqs.Enqueue(pq, 2)
	item1, _ := pqs.Dequeue(pq)
	item2, _ := pqs.Dequeue(pq)
	item3, _ := pqs.Dequeue(pq)
	assert.Equal(t, 1, item1)
	assert.Equal(t, 2, item2)
	assert.Equal(t, 3, item3)
}

func TestPeek(t *testing.T) {
	pq := pqs.New(pqs.MinFirst[int])
	pqs.Enqueue(pq, 5)
	item, ok := pqs.Peek(pq)
	assert.True(t, ok)
	assert.Equal(t, 5, item)
}

func TestPeekAfterDequeue(t *testing.T) {
	pq := pqs.New(pqs.MinFirst[int])
	pqs.Enqueue(pq, 10)
	pqs.Peek(pq)
	item, _ := pqs.Dequeue(pq)
	assert.Equal(t, 10, item)
}

func TestLen(t *testing.T) {
	pq := pqs.New(pqs.MinFirst[int])
	assert.Equal(t, 0, pqs.Len(pq))
	pqs.Enqueue(pq, 1)
	assert.Equal(t, 1, pqs.Len(pq))
	pqs.Dequeue(pq)
	assert.Equal(t, 0, pqs.Len(pq))
}

func TestClear(t *testing.T) {
	pq := pqs.New(pqs.MinFirst[int])
	pqs.Enqueue(pq, 1)
	pqs.Enqueue(pq, 2)
	pqs.Clear(pq)
	assert.Equal(t, 0, pqs.Len(pq))
	_, ok := pqs.Peek(pq)
	assert.False(t, ok)
}

func TestEmptyQueueBehavior(t *testing.T) {
	pq := pqs.New(pqs.MinFirst[int])
	_, ok1 := pqs.Peek(pq)
	_, ok2 := pqs.Dequeue(pq)
	assert.False(t, ok1)
	assert.False(t, ok2)
}

func TestSingleElementQueue(t *testing.T) {
	pq := pqs.New(pqs.MinFirst[int])
	pqs.Enqueue(pq, 7)
	peeked, ok1 := pqs.Peek(pq)
	dequeued, ok2 := pqs.Dequeue(pq)
	assert.True(t, ok1)
	assert.True(t, ok2)
	assert.Equal(t, 7, peeked)
	assert.Equal(t, 7, dequeued)
	assert.Equal(t, 0, pqs.Len(pq))
}

func TestMinFirstOrdering(t *testing.T) {
	pq := pqs.New(pqs.MinFirst[int])
	pqs.Enqueue(pq, 10)
	pqs.Enqueue(pq, 5)
	item, _ := pqs.Dequeue(pq)
	assert.Equal(t, 5, item)
}

func TestMaxFirstOrdering(t *testing.T) {
	pq := pqs.New(pqs.MaxFirst[int])
	pqs.Enqueue(pq, 10)
	pqs.Enqueue(pq, 5)
	item, _ := pqs.Dequeue(pq)
	assert.Equal(t, 10, item)
}

func Example_stringLengthPriority() {
	lengthPriority := func(x, y string) bool {
		return len(x) < len(y)
	}

	pq := pqs.New(lengthPriority)
	pqs.Enqueue(pq, "apple")
	pqs.Enqueue(pq, "kiwi")
	pqs.Enqueue(pq, "banana")

	for pqs.Len(pq) > 0 {
		item, _ := pqs.Dequeue(pq)
		fmt.Println(item)
	}
	// Output:
	// kiwi
	// apple
	// banana
}

func ExampleNew() {
	pq := pqs.New(pqs.MinFirst[int])
	pqs.Enqueue(pq, 3)
	pqs.Enqueue(pq, 1)
	item, _ := pqs.Dequeue(pq)
	fmt.Println(item)
	// Output: 1
}

func ExampleEnqueue() {
	pq := pqs.New(pqs.MinFirst[int])
	pqs.Enqueue(pq, 5)
	item, _ := pqs.Peek(pq)
	fmt.Println(item)
	// Output: 5
}

func ExampleDequeue() {
	pq := pqs.New(pqs.MinFirst[int])
	pqs.Enqueue(pq, 2)
	pqs.Enqueue(pq, 1)
	item, _ := pqs.Dequeue(pq)
	fmt.Println(item)
	// Output: 1
}

func ExamplePeek() {
	pq := pqs.New(pqs.MinFirst[int])
	pqs.Enqueue(pq, 7)
	item, _ := pqs.Peek(pq)
	fmt.Println(item)
	// Output: 7
}

func ExampleLen() {
	pq := pqs.New(pqs.MinFirst[int])
	fmt.Println(pqs.Len(pq))
	pqs.Enqueue(pq, 1)
	fmt.Println(pqs.Len(pq))
	// Output:
	// 0
	// 1
}

func ExampleClear() {
	pq := pqs.New(pqs.MinFirst[int])
	pqs.Enqueue(pq, 1)
	pqs.Clear(pq)
	fmt.Println(pqs.Len(pq))
	// Output: 0
}

func ExampleMinFirst() {
	fmt.Println(pqs.MinFirst(1, 2))
	// Output: true
}

func ExampleMaxFirst() {
	fmt.Println(pqs.MaxFirst(1, 2))
	// Output: false
}
