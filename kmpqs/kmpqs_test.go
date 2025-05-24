package kmpqs_test

import (
	"fmt"
	"testing"

	"github.com/byExist/priorityqueues/kmpqs"
	"github.com/stretchr/testify/assert"
)

type Process struct {
	PID  string
	Name string
}

func TestNew(t *testing.T) {
	q := kmpqs.New(
		kmpqs.MaxFirst[*Process, int],
		func(p *Process) string { return p.PID },
	)
	assert.NotNil(t, q)
	assert.Equal(t, 0, kmpqs.Len(q))
}

func TestClear(t *testing.T) {
	processes := map[string]*Process{
		"101": {PID: "101", Name: "nginx"},
	}
	q := kmpqs.New(
		kmpqs.MaxFirst[*Process, int],
		func(p *Process) string { return p.PID },
	)
	kmpqs.Enqueue(q, processes["101"], 1)
	kmpqs.Clear(q)
	assert.Equal(t, 0, kmpqs.Len(q))
}

func TestEnqueue(t *testing.T) {
	processes := map[string]*Process{
		"101": {PID: "101", Name: "nginx"},
	}
	q := kmpqs.New(
		kmpqs.MaxFirst[*Process, int],
		func(p *Process) string { return p.PID },
	)
	kmpqs.Enqueue(q, processes["101"], 1)
	assert.Equal(t, 1, kmpqs.Len(q))
}

func TestDequeue(t *testing.T) {
	processes := map[string]*Process{
		"101": {PID: "101", Name: "nginx"},
	}
	q := kmpqs.New(
		kmpqs.MaxFirst[*Process, int],
		func(p *Process) string { return p.PID },
	)
	kmpqs.Enqueue(q, processes["101"], 1)
	task, ok := kmpqs.Dequeue(q)
	assert.True(t, ok)
	assert.Equal(t, "101", task.PID)
	assert.Equal(t, 0, kmpqs.Len(q))
}

func TestPeek(t *testing.T) {
	processes := map[string]*Process{
		"101": {PID: "101", Name: "nginx"},
	}
	q := kmpqs.New(
		kmpqs.MaxFirst[*Process, int],
		func(p *Process) string { return p.PID },
	)
	kmpqs.Enqueue(q, processes["101"], 1)
	task, ok := kmpqs.Peek(q)
	assert.True(t, ok)
	assert.Equal(t, "101", task.PID)
	assert.Equal(t, 1, kmpqs.Len(q))
}

func TestUpdate(t *testing.T) {
	processes := map[string]*Process{
		"101": {PID: "101", Name: "nginx"},
	}
	q := kmpqs.New(
		kmpqs.MaxFirst[*Process, int],
		func(p *Process) string { return p.PID },
	)
	kmpqs.Enqueue(q, processes["101"], 1)
	processes["101"].Name = "nginx-updated"
	kmpqs.Update(q, processes["101"], 5)
	task, _ := kmpqs.Peek(q)
	assert.Equal(t, "nginx-updated", task.Name)
}

func TestDelete(t *testing.T) {
	processes := map[string]*Process{
		"101": {PID: "101", Name: "nginx"},
	}
	q := kmpqs.New(
		kmpqs.MaxFirst[*Process, int],
		func(p *Process) string { return p.PID },
	)
	kmpqs.Enqueue(q, processes["101"], 1)
	kmpqs.Delete(q, processes["101"])
	assert.Equal(t, 0, kmpqs.Len(q))
}

func TestLen(t *testing.T) {
	processes := map[string]*Process{
		"101": {PID: "101", Name: "nginx"},
	}
	q := kmpqs.New(
		kmpqs.MaxFirst[*Process, int],
		func(p *Process) string { return p.PID },
	)
	assert.Equal(t, 0, kmpqs.Len(q))
	kmpqs.Enqueue(q, processes["101"], 1)
	assert.Equal(t, 1, kmpqs.Len(q))
}

func TestContains(t *testing.T) {
	processes := map[string]*Process{
		"101": {PID: "101", Name: "nginx"},
		"102": {PID: "102", Name: "postgres"},
	}
	q := kmpqs.New(
		kmpqs.MaxFirst[*Process, int],
		func(p *Process) string { return p.PID },
	)
	kmpqs.Enqueue(q, processes["101"], 1)
	assert.True(t, kmpqs.Contains(q, processes["101"]))
	assert.False(t, kmpqs.Contains(q, processes["102"]))
}

func ExampleNew() {
	type Process struct {
		PID  string
		Name string
	}
	q := kmpqs.New(
		kmpqs.MaxFirst[*Process, int],
		func(p *Process) string { return p.PID },
	)
	fmt.Println(kmpqs.Len(q))
	// Output: 0
}

func ExampleClear() {
	type Process struct {
		PID  string
		Name string
	}
	processes := map[string]*Process{
		"101": {PID: "101", Name: "nginx"},
	}
	q := kmpqs.New(
		kmpqs.MaxFirst[*Process, int],
		func(p *Process) string { return p.PID },
	)
	kmpqs.Enqueue(q, processes["101"], 1)
	kmpqs.Clear(q)
	fmt.Println(kmpqs.Len(q))
	// Output: 0
}

func ExampleEnqueue() {
	type Process struct {
		PID  string
		Name string
	}
	processes := map[string]*Process{
		"101": {PID: "101", Name: "nginx"},
	}
	q := kmpqs.New(
		kmpqs.MaxFirst[*Process, int],
		func(p *Process) string { return p.PID },
	)
	kmpqs.Enqueue(q, processes["101"], 1)
	fmt.Println(kmpqs.Len(q))
	// Output: 1
}

func ExampleDequeue() {
	type Process struct {
		PID  string
		Name string
	}
	processes := map[string]*Process{
		"101": {PID: "101", Name: "nginx"},
	}
	q := kmpqs.New(
		kmpqs.MaxFirst[*Process, int],
		func(p *Process) string { return p.PID },
	)
	kmpqs.Enqueue(q, processes["101"], 1)
	task, ok := kmpqs.Dequeue(q)
	fmt.Println(ok, task.PID)
	// Output: true 101
}

func ExamplePeek() {
	type Process struct {
		PID  string
		Name string
	}
	processes := map[string]*Process{
		"101": {PID: "101", Name: "nginx"},
	}
	q := kmpqs.New(
		kmpqs.MaxFirst[*Process, int],
		func(p *Process) string { return p.PID },
	)
	kmpqs.Enqueue(q, processes["101"], 1)
	task, ok := kmpqs.Peek(q)
	fmt.Println(ok, task.PID)
	// Output: true 101
}

func ExampleUpdate() {
	type Process struct {
		PID  string
		Name string
	}
	processes := map[string]*Process{
		"101": {PID: "101", Name: "nginx"},
	}
	q := kmpqs.New(
		kmpqs.MaxFirst[*Process, int],
		func(p *Process) string { return p.PID },
	)
	kmpqs.Enqueue(q, processes["101"], 1)
	kmpqs.Update(q, processes["101"], 5)
	task, _ := kmpqs.Peek(q)
	fmt.Println(task.Name)
	fmt.Println(kmpqs.Len(q))
	// Output:
	// nginx
	// 1
}

func ExampleDelete() {
	type Process struct {
		PID  string
		Name string
	}
	processes := map[string]*Process{
		"101": {PID: "101", Name: "nginx"},
	}
	q := kmpqs.New(
		kmpqs.MaxFirst[*Process, int],
		func(p *Process) string { return p.PID },
	)
	kmpqs.Enqueue(q, processes["101"], 1)
	kmpqs.Delete(q, processes["101"])
	fmt.Println(kmpqs.Len(q))
	// Output: 0
}

func ExampleLen() {
	type Process struct {
		PID  string
		Name string
	}
	processes := map[string]*Process{
		"101": {PID: "101", Name: "nginx"},
	}
	q := kmpqs.New(
		kmpqs.MaxFirst[*Process, int],
		func(p *Process) string { return p.PID },
	)
	fmt.Println(kmpqs.Len(q))
	kmpqs.Enqueue(q, processes["101"], 1)
	fmt.Println(kmpqs.Len(q))
	// Output:
	// 0
	// 1
}

func ExampleContains() {
	type Process struct {
		PID  string
		Name string
	}
	processes := map[string]*Process{
		"101": {PID: "101", Name: "nginx"},
		"102": {PID: "102", Name: "postgres"},
	}
	q := kmpqs.New(
		kmpqs.MaxFirst[*Process, int],
		func(p *Process) string { return p.PID },
	)
	kmpqs.Enqueue(q, processes["101"], 1)
	fmt.Println(kmpqs.Contains(q, processes["101"]))
	fmt.Println(kmpqs.Contains(q, processes["102"]))
	// Output:
	// true
	// false
}
