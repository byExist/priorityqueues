package priorityqueues_test

import (
	"fmt"
	"slices"
	"testing"

	pqs "github.com/byExist/priorityqueues"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	pq := pqs.New(func(x int) int { return x })
	require.NotNil(t, pq)
	require.Equal(t, 0, pqs.Len(pq))
}

func TestCollect(t *testing.T) {
	emptySeq := func(yield func(int) bool) {}
	pq := pqs.Collect(emptySeq, func(x int) int { return x })
	assert.NotNil(t, pq)
	assert.Equal(t, 0, pqs.Len(pq))

	values := []int{3, 1, 4, 2}
	pq = pqs.Collect(slices.Values(values), func(x int) int { return x })
	assert.Equal(t, len(values), pqs.Len(pq))

	prev, ok := pqs.Peek(pq)
	require.True(t, ok)
	for range values {
		next, ok := pqs.Dequeue(pq)
		require.True(t, ok)
		assert.GreaterOrEqual(t, prev, next)
		prev = next
	}
}

func TestClone(t *testing.T) {
	pq := pqs.New(func(x int) int { return x })
	pqs.Enqueue(pq, 10)
	pqs.Enqueue(pq, 20)

	clone := pqs.Clone(pq)

	pqs.Enqueue(clone, 30)
	assert.NotEqual(t, pqs.Len(pq), pqs.Len(clone))
}

func TestEnqueue(t *testing.T) {
	pq := pqs.New(func(x int) int { return x })
	pqs.Enqueue(pq, 5)
	assert.Equal(t, 1, pqs.Len(pq))

	pqs.Enqueue(pq, 10)
	pqs.Enqueue(pq, 1)

	top, ok := pqs.Peek(pq)
	require.True(t, ok)
	assert.Equal(t, 10, top)
}

func TestDequeue(t *testing.T) {
	pq := pqs.New(func(x int) int { return x })

	_, ok := pqs.Dequeue(pq)
	assert.False(t, ok)

	pqs.Enqueue(pq, 3)
	pqs.Enqueue(pq, 7)

	val, ok := pqs.Dequeue(pq)
	require.True(t, ok)
	assert.Equal(t, 7, val)

	val, ok = pqs.Dequeue(pq)
	require.True(t, ok)
	assert.Equal(t, 3, val)

	_, ok = pqs.Dequeue(pq)
	assert.False(t, ok)
}

func TestPeek(t *testing.T) {
	pq := pqs.New(func(x int) int { return x })

	_, ok := pqs.Peek(pq)
	assert.False(t, ok)

	pqs.Enqueue(pq, 8)
	val, ok := pqs.Peek(pq)
	require.True(t, ok)
	assert.Equal(t, 8, val)

	assert.Equal(t, 1, pqs.Len(pq))
}

func TestLen(t *testing.T) {
	pq := pqs.New(func(x int) int { return x })
	assert.Equal(t, 0, pqs.Len(pq))

	pqs.Enqueue(pq, 1)
	assert.Equal(t, 1, pqs.Len(pq))

	pqs.Dequeue(pq)
	assert.Equal(t, 0, pqs.Len(pq))
}

func TestValues(t *testing.T) {
	pq := pqs.New(func(x int) int { return x })

	count := 0
	for range pqs.Values(pq) {
		count++
	}
	assert.Equal(t, 0, count)

	pqs.Enqueue(pq, 2)
	pqs.Enqueue(pq, 5)
	pqs.Enqueue(pq, 1)

	expected := []int{5, 2, 1}
	results := slices.Collect(pqs.Values(pq))
	assert.ElementsMatch(t, expected, results)
}

func TestClear(t *testing.T) {
	pq := pqs.New(func(x int) int { return x })

	pqs.Enqueue(pq, 1)
	pqs.Enqueue(pq, 2)
	assert.NotEqual(t, 0, pqs.Len(pq))

	pqs.Clear(pq)
	assert.Equal(t, 0, pqs.Len(pq))

	pqs.Enqueue(pq, 3)
	val, ok := pqs.Peek(pq)
	require.True(t, ok)
	assert.Equal(t, 3, val)
}

func TestVariousTypes(t *testing.T) {
	longer := func(s string) int { return len(s) }
	strPQ := pqs.New(longer)
	pqs.Enqueue(strPQ, "short")
	pqs.Enqueue(strPQ, "longer")
	pqs.Enqueue(strPQ, "longest")
	val, ok := pqs.Dequeue(strPQ)
	require.True(t, ok)
	assert.Equal(t, "longest", val)

	larger := func(f float64) float64 { return f }
	floatPQ := pqs.New(larger)
	pqs.Enqueue(floatPQ, 1.5)
	pqs.Enqueue(floatPQ, 2.3)
	pqs.Enqueue(floatPQ, 0.7)
	valF, ok := pqs.Dequeue(floatPQ)
	require.True(t, ok)
	assert.Equal(t, 2.3, valF)

	type person struct {
		name string
		age  int
	}
	older := func(p person) int { return p.age }
	peoplePQ := pqs.New(older)
	pqs.Enqueue(peoplePQ, person{"Alice", 30})
	pqs.Enqueue(peoplePQ, person{"Bob", 25})
	pqs.Enqueue(peoplePQ, person{"Charlie", 35})
	valP, ok := pqs.Dequeue(peoplePQ)
	require.True(t, ok)
	assert.Equal(t, person{"Charlie", 35}, valP)
}

// ExampleNew demonstrates usage of New and Enqueue/Dequeue.
func ExampleNew() {
	larger := func(x int) int { return x }
	pq := pqs.New(larger)

	pqs.Enqueue(pq, 10)
	pqs.Enqueue(pq, 30)
	pqs.Enqueue(pq, 20)

	for pqs.Len(pq) > 0 {
		v, _ := pqs.Dequeue(pq)
		fmt.Println(v)
	}
	// Output:
	// 30
	// 20
	// 10
}

func ExampleNew_struct() {
	type person struct {
		name string
		age  int
	}

	older := func(p person) int { return p.age }
	pq := pqs.New(older)
	pqs.Enqueue(pq, person{"Alice", 30})
	pqs.Enqueue(pq, person{"Bob", 25})
	pqs.Enqueue(pq, person{"Charlie", 35})

	p, _ := pqs.Dequeue(pq)
	fmt.Println(p.name, p.age)
	// Output:
	// Charlie 35
}

// ExampleCollect demonstrates usage of Collect with a generator.
func ExampleCollect() {
	seq := slices.Values([]int{1, 3, 2})
	pq := pqs.Collect(seq, func(x int) int { return x })
	v, _ := pqs.Dequeue(pq)
	fmt.Println(v)
	// Output:
	// 3
}

// ExampleClone shows that Clone creates an independent copy.
func ExampleClone() {
	larger := func(x int) int { return x }
	pq := pqs.New(larger)
	pqs.Enqueue(pq, 5)

	clone := pqs.Clone(pq)
	pqs.Enqueue(clone, 10)

	s1 := slices.Collect(pqs.Values(pq))
	s2 := slices.Collect(pqs.Values(clone))

	fmt.Println(s1)
	fmt.Println(s2)
	// Output:
	// [5]
	// [10 5]
}

// ExampleEnqueue demonstrates usage of Enqueue and Peek.
func ExampleEnqueue() {
	pq := pqs.New(func(x int) int { return x })
	pqs.Enqueue(pq, 1)
	pqs.Enqueue(pq, 3)
	pqs.Enqueue(pq, 2)
	v, _ := pqs.Peek(pq)
	fmt.Println(v)
	// Output:
	// 3
}

// ExampleDequeue demonstrates usage of Dequeue.
func ExampleDequeue() {
	pq := pqs.New(func(x int) int { return x })
	pqs.Enqueue(pq, 4)
	pqs.Enqueue(pq, 2)
	v, _ := pqs.Dequeue(pq)
	fmt.Println(v)
	// Output:
	// 4
}

// ExamplePeek demonstrates usage of Peek.
func ExamplePeek() {
	pq := pqs.New(func(x int) int { return x })
	pqs.Enqueue(pq, 7)
	v, _ := pqs.Peek(pq)
	fmt.Println(v)
	// Output:
	// 7
}

// ExampleLen demonstrates usage of Len.
func ExampleLen() {
	pq := pqs.New(func(x int) int { return x })
	fmt.Println(pqs.Len(pq))
	pqs.Enqueue(pq, 5)
	fmt.Println(pqs.Len(pq))
	// Output:
	// 0
	// 1
}

// ExampleValues demonstrates usage of Values.
func ExampleValues() {
	pq := pqs.New(func(x int) int { return x })
	pqs.Enqueue(pq, 10)
	pqs.Enqueue(pq, 30)
	pqs.Enqueue(pq, 20)
	for v := range pqs.Values(pq) {
		fmt.Println(v)
	}
	// Output:
	// 30
	// 10
	// 20
}

// ExampleClear demonstrates usage of Clear.
func ExampleClear() {
	pq := pqs.New(func(x int) int { return x })
	pqs.Enqueue(pq, 1)
	pqs.Enqueue(pq, 2)
	pqs.Clear(pq)
	fmt.Println(pqs.Len(pq))
	// Output:
	// 0
}
