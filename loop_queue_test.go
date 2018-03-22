// Package ds provides ...
package ds

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	queue := new(LoopQueue)

	if !queue.IsEmpty() {
		t.Errorf("queue.IsEmpty() == %t, want %t", queue.IsEmpty(), true)
	}

	queue.Insert(1)
	if queue.IsEmpty() {
		t.Errorf("queue.IsEmpty() == %t, want %t", queue.IsEmpty(), false)
	}

	queue.Remove()
	if !queue.IsEmpty() {
		t.Errorf("queue.IsEmpty() == %t, want %t", queue.IsEmpty(), true)
	}
}

func TestIsFull(t *testing.T) {
	queue := new(LoopQueue)

	for i := 0; i < 10; i++ {
		queue.Insert(i)
	}

	if !queue.IsFull() {
		t.Errorf("queue.IsFull() == %t, want %t", queue.IsFull(), true)
	}

	queue.Remove()
	if queue.IsFull() {
		t.Errorf("queue.IsFull() == %t, want %t", queue.IsFull(), false)
	}

	queue.Insert(11)
	if !queue.IsFull() {
		t.Errorf("queue.IsFull() == %t, want %t", queue.IsFull(), true)
	}
}

func TestLength(t *testing.T) {
	queue := new(LoopQueue)

	if queue.Length() != 0 {
		t.Errorf("queue.Length() == %v, want %v", queue.Length(), 0)
	}

	for i := 0; i < 20; i++ {
		queue.Insert(i)
	}
	if queue.Length() != 9 {
		t.Errorf("queue.Length() == %v, want %v", queue.Length(), 9)
	}
}
