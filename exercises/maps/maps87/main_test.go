// maps87
// Make the tests pass!

// I AM NOT DONE
//
// Counter накапливает счётчики. Для нового значения Counter{} запись паникует.
// Тренирует: нулевое значение map — nil, запись в неё вызывает панику.
// Сложность: hard
package main_test

import "testing"

type Counter struct {
	counts map[string]int
}

func (c *Counter) Add(key string) {
	c.counts[key]++
}

func (c *Counter) Get(key string) int { return c.counts[key] }

func TestCounter(t *testing.T) {
	var c Counter
	if c.Get("x") != 0 {
		t.Errorf("Get on empty counter should be 0")
	}
	c.Add("x")
	c.Add("x")
	if c.Get("x") != 2 {
		t.Errorf("Get(x) = %d, want 2", c.Get("x"))
	}
}
