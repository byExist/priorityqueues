# pqs [![GoDoc](https://pkg.go.dev/badge/github.com/byExist/priorityqueues/pqs.svg)](https://pkg.go.dev/github.com/byExist/priorityqueues/pqs) [![Go Report Card](https://goreportcard.com/badge/github.com/byExist/priorityqueues)](https://goreportcard.com/report/github.com/byExist/priorityqueues)

A minimal, generic priority queue in Go using ordered types.

The `pqs` package provides the simplest form of a priority queue for values that are inherently comparable (e.g., `int`, `float64`, `string`). It does not support key-based lookups or custom structures — items are their own priority.

---

## ✨ Features

- ✅ Minimal: single-type input, no struct wrapping
- ✅ Generic: works with any `cmp.Ordered` type
- ✅ Custom comparator: control min/max or custom logic
- ❌ No stability guarantees (insertion order not preserved for equal priority)
- ❌ No key support or item updates

---

## 🧱 Example

```go
package main

import (
	"fmt"
	"github.com/byExist/priorityqueues/pqs"
)

func main() {
	q := pqs.New(pqs.MinFirst[int])
	pqs.Enqueue(q, 3)
	pqs.Enqueue(q, 1)
	pqs.Enqueue(q, 2)

	for pqs.Len(q) > 0 {
		item, _ := pqs.Dequeue(q)
		fmt.Println(item)
	}
}

// Output:
// 1
// 2
// 3
```

---

## 📚 Use When

- You need a fast, generic min/max heap
- The item itself is the priority
- You don’t need key lookups, updates, or custom types

---

## 🚫 Avoid If

- You need to update item priority (`kmpqs`, `kpqs`)
- You need key-based identity (`kpqs`, `kmpqs`)
- You want stable ordering among equal-priority items (`mpqs`, `kpqs`)

---

## 🔍 Alternatives

| Package | Use case |
|---------|----------|
| `mpqs`  | custom structs with external priority |
| `kpqs`  | key + internal priority field |
| `kmpqs` | key + external priority control |
