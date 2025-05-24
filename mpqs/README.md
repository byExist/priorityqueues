# mpqs [![GoDoc](https://pkg.go.dev/badge/github.com/byExist/priorityqueues/mpqs.svg)](https://pkg.go.dev/github.com/byExist/priorityqueues/mpqs) [![Go Report Card](https://goreportcard.com/badge/github.com/byExist/priorityqueues)](https://goreportcard.com/report/github.com/byExist/priorityqueues)

A generic, manually prioritized queue with stability.

The `mpqs` package provides a priority queue implementation where the user explicitly provides the priority value for each item. It supports custom types, stable ordering (tie-breaking by insertion order), and comparator injection.

---

## ✨ Features

- ✅ Generic over item type (`T`) and priority type (`P`)
- ✅ External priority injection per enqueue
- ✅ Stable ordering for equal priority values
- ✅ Custom comparator support (min, max, stable variants)
- ❌ No key-based lookup or update support

---

## 🧱 Example

```go
package main

import (
	"fmt"
	"github.com/byExist/priorityqueues/mpqs"
)

type Task struct {
	ID   string
	Desc string
}

func main() {
	q := mpqs.New(mpqs.StableMinFirst[*Task, int])

	tasks := []*Task{
		{ID: "t1", Desc: "send email"},
		{ID: "t2", Desc: "render image"},
		{ID: "t3", Desc: "generate report"},
	}

	// External priorities (lower means higher priority)
	mpqs.Enqueue(q, tasks[0], 3) // email
	mpqs.Enqueue(q, tasks[1], 2) // image
	mpqs.Enqueue(q, tasks[2], 1) // report

	for mpqs.Len(q) > 0 {
		task, _ := mpqs.Dequeue(q)
		fmt.Println(task.ID, ":", task.Desc)
	}
}

// Output:
// t3 : generate report
// t2 : render image
// t1 : send email
```

---

## 📚 Use When

- You want full control over priority values at enqueue time
- You need stable ordering among equal-priority items
- You don't need key-based access or updates

---

## 🚫 Avoid If

- You need to identify/update/delete items by key → use `kpqs` or `kmpqs`
- Your item inherently contains its own priority → use `kpqs`
- You want a minimal value-only queue → use `pqs`

---

## 🔍 Comparator Options

Use one of the following with `New(...)`:

- `mpqs.MinFirst[T, P]`
- `mpqs.MaxFirst[T, P]`
- `mpqs.StableMinFirst[T, P]`
- `mpqs.StableMaxFirst[T, P]`

You can also provide a custom comparator function.
