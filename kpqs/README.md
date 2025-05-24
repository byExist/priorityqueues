# kpqs [![GoDoc](https://pkg.go.dev/badge/github.com/byExist/priorityqueues/kpqs.svg)](https://pkg.go.dev/github.com/byExist/priorityqueues/kpqs) [![Go Report Card](https://goreportcard.com/badge/github.com/byExist/priorityqueues)](https://goreportcard.com/report/github.com/byExist/priorityqueues)

A generic keyed priority queue where priority is derived from the item itself.

The `kpqs` package provides a stable, key-addressable priority queue for arbitrary types. Each item is assigned a priority via a user-provided function. Items are inserted with a key (derived from the item), and duplicates are overwritten.

---

## ✨ Features

- ✅ Key-based access (update, delete, contains)
- ✅ Priority derived from item field or logic
- ✅ Stable ordering: earlier enqueued wins on tie
- ✅ Comparator injection (`MinFirst`, `MaxFirst`, etc.)
- ❌ No external priority control at enqueue time

---

## 🧱 Example

```go
package main

import (
	"fmt"
	"github.com/byExist/priorityqueues/kpqs"
)

type Task struct {
	ID       string
	Priority int
}

func main() {
	q := kpqs.New(
		kpqs.StableMinFirst[*Task, int],
		func(t *Task) string { return t.ID },
		func(t *Task) int { return t.Priority },
	)

	kpqs.Enqueue(q, &Task{ID: "a", Priority: 3})
	kpqs.Enqueue(q, &Task{ID: "b", Priority: 1})
	kpqs.Enqueue(q, &Task{ID: "c", Priority: 2})

	for kpqs.Len(q) > 0 {
		task, _ := kpqs.Dequeue(q)
		fmt.Println(task.ID)
	}
}

// Output:
// b
// c
// a
```

---

## 📚 Use When

- You want a keyed queue (e.g., `map[ID]*Task`)
- Each item knows its own priority
- You need stable, deterministic dequeue order

---

## 🚫 Avoid If

- You want to provide priority externally → use `kmpqs`
- You don’t need keys or updates → use `mpqs` or `pqs`

---

## 🔍 Comparator Options

Use one of the following with `New(...)`:

- `kpqs.MinFirst[T, P]`
- `kpqs.MaxFirst[T, P]`
- `kpqs.StableMinFirst[T, P]`
- `kpqs.StableMaxFirst[T, P]`

You can also provide a custom comparator function.
