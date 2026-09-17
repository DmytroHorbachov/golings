// concurrent99
// Make the tests pass!

// I AM NOT DONE
//
// Get читает под RLock, но освобождает не тот замок, и программа падает.
// Тренирует: парность методов RWMutex.
// Сложность: easy
package main_test

import (
	"sync"
	"testing"
)

type Registry struct {
	mu    sync.RWMutex
	items map[string]int
}

func (r *Registry) Get(k string) int {
	r.mu.RLock()
	defer r.mu.Unlock()
	return r.items[k]
}

func TestRegistry(t *testing.T) {
	r := &Registry{items: map[string]int{"a": 1}}
	if r.Get("a") != 1 || r.Get("a") != 1 {
		t.Errorf("Get failed")
	}
}
