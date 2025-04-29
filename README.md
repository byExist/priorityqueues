# priorityqueues [![GoDoc](https://pkg.go.dev/badge/github.com/byExist/priorityqueues.svg)](https://pkg.go.dev/github.com/byExist/priorityqueues) [![Go Report Card](https://goreportcard.com/badge/github.com/byExist/priorityqueues)](https://goreportcard.com/report/github.com/byExist/priorityqueues)

## What is "priorityqueues"?

**priorityqueues** is a lightweight and generic priority queue package for Go. It supports custom priority functions and works with any comparable types. Internally, it uses a binary heap to ensure efficient enqueue and dequeue operations based on priority.

## Features

- Supports generic types
- Works with custom priority functions
- Efficient Enqueue/Dequeue using `container/heap`
- Cloneable priority queues
- Clear and reuse queues
- Iterable with `iter.Seq`

## Installation

To install, run:

```bash
go get github.com/byExist/priorityqueues
```

## Quick Start

```go
package main

import (
	"fmt"

	pqs "github.com/byExist/priorityqueues"
)

type task struct {
	name     string
	priority int
}

func main() {
	// Create a priority queue with custom priority function
	priority := func(t task) int {
		return t.priority
	}
	pq := pqs.New(priority)

	// Add tasks to the queue
	pqs.Enqueue(pq, task{"write docs", 2})
	pqs.Enqueue(pq, task{"fix bug", 5})
	pqs.Enqueue(pq, task{"implement feature", 3})

	// Process tasks in order of priority
	for {
		t, ok := pqs.Dequeue(pq)
		if !ok {
			break
		}
		fmt.Println("Processing:", t.name)
	}
}
```

## Usage

The `priorityqueues` package allows you to create and manage efficient, type-safe priority queues using any Go type. You can define your own priority function and use generic functions to enqueue, dequeue, peek, clear, clone, and iterate through the queue.

## API Overview

### Constructors

- `New[T, P](priorityFunc func(T) P) *PriorityQueue[T, P]`
- `FromSeq[T, P](seq iter.Seq[T], priorityFunc func(T) P) *PriorityQueue[T, P]`

### Core Functions

- `Clone(pq *PriorityQueue[T, P]) *PriorityQueue[T, P]`
- `Enqueue(pq *PriorityQueue[T, P], item T)`
- `Dequeue(pq *PriorityQueue[T, P]) (T, bool)`
- `Peek(pq *PriorityQueue[T, P]) (T, bool)`
- `Clear(pq *PriorityQueue[T, P])`
- `Values(pq *PriorityQueue[T, P]) iter.Seq[T]`
- `Len(pq *PriorityQueue[T, P]) int`

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.