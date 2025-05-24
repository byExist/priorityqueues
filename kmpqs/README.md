# kmpqs [![GoDoc](https://pkg.go.dev/badge/github.com/byExist/priorityqueues/kmpqs.svg)](https://pkg.go.dev/github.com/byExist/priorityqueues/kmpqs) [![Go Report Card](https://goreportcard.com/badge/github.com/byExist/priorityqueues)](https://goreportcard.com/report/github.com/byExist/priorityqueues)

A generic, keyed priority queue with external priority control.

The `kmpqs` package provides a stable, key-addressable priority queue where the priority is explicitly provided at enqueue/update time. It is suitable for task schedulers, job queues, or systems where the item does not inherently contain its own priority.

---

## ✨ Features

- ✅ Key-based lookup and overwrite
- ✅ External priority injection (`Enqueue(item, prio)`)
- ✅ Stable ordering (insertion order respected on tie)
- ✅ Supports `Update`, `Delete`, `Contains`
- ✅ Custom comparator support (min, max, stable)
- ❌ Priority is not extracted from item

---

## 🧱 Example

```go
package main

import (
	"fmt"
	"github.com/byExist/priorityqueues/kmpqs"
)

type Process struct {
	PID  string
	Name string
}

func main() {
	q := kmpqs.New(
		kmpqs.StableMinFirst[*Process, int],
		func(p *Process) string { return p.PID },
	)

	processes := map[string]*Process{
		"101": {PID: "101", Name: "nginx"},
		"102": {PID: "102", Name: "postgres"},
	}

	kmpqs.Enqueue(q, processes["101"], 5)
	kmpqs.Enqueue(q, processes["102"], 3)

	// Change priority later
	kmpqs.Update(q, processes["101"], 1)

	for kmpqs.Len(q) > 0 {
		p, _ := kmpqs.Dequeue(q)
		fmt.Println(p.Name)
	}
}

// Output:
// nginx
// postgres
```

---

## 📚 Use When

- You want to manage queue item priority **externally**
- You want to update or replace queued items by key
- You want deterministic ordering (FIFO on same priority)

---

## 🚫 Avoid If

- Your item already has its own priority field → use `kpqs`
- You don’t need key-based identity → use `mpqs` or `pqs`

---

## 🔍 Comparator Options

Use one of the following with `New(...)`:

- `kmpqs.MinFirst[T, P]`
- `kmpqs.MaxFirst[T, P]`
- `kmpqs.StableMinFirst[T, P]`
- `kmpqs.StableMaxFirst[T, P]`

You can also provide a custom comparator function.
