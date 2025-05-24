# PriorityQueues

**A collection of generic and modular priority queue implementations in Go.**

This repository provides four distinct priority queue strategies, each targeting a specific use case or trade-off between simplicity, flexibility, and key-based access.

---

## 📦 Packages Overview

| Package  | Key Support | Priority Injection | Stability | Typical Use Case |
|----------|-------------|--------------------|-----------|------------------|
| `pqs`    | ❌           | ❌ (item is prio)   | ❌         | Simple int/float/string queue |
| `mpqs`   | ❌           | ✅                  | ✅         | Manual prio for structs |
| `kpqs`   | ✅           | ❌ (prio from item) | ✅         | Tasks with embedded priority |
| `kmpqs`  | ✅           | ✅                  | ✅         | Schedulers, process queues |

Each package is self-contained and independently tested.

---

## ✨ Why multiple queues?

Different use cases require different queue behavior:

- **Minimal overhead**? Use `pqs`.
- **Custom priority from a field**? Use `kpqs`.
- **Keyed access with externally determined priority**? Use `kmpqs`.
- **Just need control over the comparator**? Use `mpqs`.

---

## 📂 Structure

```
priorityqueues/
├── kmpqs/  // keyed + manual prio
├── kpqs/   // keyed + prio from item
├── mpqs/   // manual prio only
├── pqs/    // basic queue
```

Each directory contains:

- `*.go` – package implementation  
- `*_test.go` – unit tests and examples  
- `README.md` – (optional) package-specific usage and design notes

---

## ✅ Features

- Go 1.18+ generic support
- `container/heap` under the hood
- Stable priority resolution with tie-breaking by insertion order
- Optional key-based lookup (`kmpqs`, `kpqs`)
- Custom comparator functions

---

## 🔍 Getting Started

Install with:

```sh
go get github.com/byExist/priorityqueues
```

Import the package you need:

```go
import "github.com/byExist/priorityqueues/kmpqs"
```

Then initialize and use:

```go
q := kmpqs.New(
	kmpqs.StableMinFirst[*Process, int],
	func(p *Process) string { return p.PID },
)
kmpqs.Enqueue(q, &Process{PID: "123", Name: "nginx"}, 5)
```

---

## 🧪 Testing

Each package includes full test coverage and practical `Example` functions.

Run all tests with:

```sh
go test ./...
```

---

## 📚 See Also

- [Go container/heap](https://pkg.go.dev/container/heap)
- [Go Generics](https://go.dev/doc/tutorial/generics)

---

## 📄 License

MIT License. See [LICENSE](./LICENSE) for details.