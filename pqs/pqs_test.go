package pqs_test

import (
	"fmt"
	"testing"

	"github.com/byExist/priorityqueues/pqs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMinFirst(t *testing.T) {
	require.True(t, pqs.MinFirst(1, 2))
	require.False(t, pqs.MinFirst(2, 1))
}

func TestMaxFirst(t *testing.T) {
	require.True(t, pqs.MaxFirst(3, 2))
	require.False(t, pqs.MaxFirst(1, 4))
}

func TestNew(t *testing.T) {
	pq := pqs.New(pqs.MinFirst[int])
	require.NotNil(t, pq)
	assert.Equal(t, 0, pqs.Len(pq))
}

func TestClear(t *testing.T) {
	pq := pqs.New(pqs.MinFirst[int])
	pqs.Enqueue(pq, 1)
	pqs.Enqueue(pq, 2)
	pqs.Clear(pq)
	assert.Equal(t, 0, pqs.Len(pq))
}

func TestEnqueue(t *testing.T) {
	pq := pqs.New(pqs.MinFirst[int])
	pqs.Enqueue(pq, 10)
	assert.Equal(t, 1, pqs.Len(pq))
	top, ok := pqs.Peek(pq)
	require.True(t, ok)
	assert.Equal(t, 10, top)
}

func TestDequeue(t *testing.T) {
	pq := pqs.New(pqs.MinFirst[int])
	pqs.Enqueue(pq, 3)
	pqs.Enqueue(pq, 1)
	pqs.Enqueue(pq, 2)

	v, ok := pqs.Dequeue(pq)
	require.True(t, ok)
	assert.Equal(t, 1, v)

	v, ok = pqs.Dequeue(pq)
	require.True(t, ok)
	assert.Equal(t, 2, v)

	v, ok = pqs.Dequeue(pq)
	require.True(t, ok)
	assert.Equal(t, 3, v)
}

func TestDequeue_Empty(t *testing.T) {
	pq := pqs.New(pqs.MinFirst[int])
	_, ok := pqs.Dequeue(pq)
	assert.False(t, ok)
}

func TestPeek(t *testing.T) {
	pq := pqs.New(pqs.MinFirst[int])
	pqs.Enqueue(pq, 5)
	pqs.Enqueue(pq, 1)
	top, ok := pqs.Peek(pq)
	require.True(t, ok)
	assert.Equal(t, 1, top)
	assert.Equal(t, 2, pqs.Len(pq)) // Peek should not remove
}

func TestPeek_Empty(t *testing.T) {
	pq := pqs.New(pqs.MinFirst[int])
	_, ok := pqs.Peek(pq)
	assert.False(t, ok)
}

func TestLen(t *testing.T) {
	pq := pqs.New(pqs.MinFirst[int])
	assert.Equal(t, 0, pqs.Len(pq))
	pqs.Enqueue(pq, 10)
	assert.Equal(t, 1, pqs.Len(pq))
}

func Example_minLengthString() {
	less := func(a, b string) bool {
		return len(a) < len(b)
	}

	pq := pqs.New(less)
	pqs.Enqueue(pq, "banana")
	pqs.Enqueue(pq, "kiwi")
	pqs.Enqueue(pq, "apple")
	pqs.Enqueue(pq, "fig")

	for pqs.Len(pq) > 0 {
		v, _ := pqs.Dequeue(pq)
		fmt.Println(v)
	}
	// Output:
	// fig
	// kiwi
	// apple
	// banana
}

func ExampleMinFirst() {
	pq := pqs.New(pqs.MinFirst[int])
	pqs.Enqueue(pq, 5)
	pqs.Enqueue(pq, 2)
	pqs.Enqueue(pq, 3)

	for pqs.Len(pq) > 0 {
		v, _ := pqs.Dequeue(pq)
		fmt.Println(v)
	}
	// Output:
	// 2
	// 3
	// 5
}

func ExampleMaxFirst() {
	pq := pqs.New(pqs.MaxFirst[int])
	pqs.Enqueue(pq, 1)
	pqs.Enqueue(pq, 4)
	pqs.Enqueue(pq, 2)

	for pqs.Len(pq) > 0 {
		v, _ := pqs.Dequeue(pq)
		fmt.Println(v)
	}
	// Output:
	// 4
	// 2
	// 1
}

func ExampleNew() {
	pq := pqs.New(pqs.MinFirst[int])
	pqs.Enqueue(pq, 3)
	pqs.Enqueue(pq, 1)
	for pqs.Len(pq) > 0 {
		v, _ := pqs.Dequeue(pq)
		fmt.Println(v)
	}
	// Output:
	// 1
	// 3
}

func ExampleClear() {
	pq := pqs.New(pqs.MinFirst[int])
	pqs.Enqueue(pq, 10)
	pqs.Clear(pq)
	fmt.Println("len after clear:", pqs.Len(pq))
	// Output:
	// len after clear: 0
}

func ExampleEnqueue() {
	pq := pqs.New(pqs.MinFirst[int])
	pqs.Enqueue(pq, 42)
	v, ok := pqs.Peek(pq)
	if ok {
		fmt.Println("enqueued:", v)
	}
	// Output:
	// enqueued: 42
}

func ExampleDequeue() {
	pq := pqs.New(pqs.MinFirst[int])
	pqs.Enqueue(pq, 3)
	pqs.Enqueue(pq, 1)
	v, ok := pqs.Dequeue(pq)
	if ok {
		fmt.Println("dequeued:", v)
	}
	// Output:
	// dequeued: 1
}

func ExamplePeek() {
	pq := pqs.New(pqs.MinFirst[int])
	pqs.Enqueue(pq, 7)
	pqs.Enqueue(pq, 3)

	v, ok := pqs.Peek(pq)
	if ok {
		fmt.Println("peek:", v)
	}
	fmt.Println("len:", pqs.Len(pq))
	// Output:
	// peek: 3
	// len: 2
}

func ExampleLen() {
	pq := pqs.New(pqs.MinFirst[int])
	fmt.Println("len:", pqs.Len(pq))
	pqs.Enqueue(pq, 7)
	fmt.Println("len after enqueue:", pqs.Len(pq))
	// Output:
	// len: 0
	// len after enqueue: 1
}
