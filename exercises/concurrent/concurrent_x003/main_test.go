// concurrent_x003: Буферизованный канал
// Make the tests pass!
// I AM NOT DONE
//
// store отправляет значение в канал и тут же его читает в той же горутине.
// Без буфера отправка блокируется навсегда.
// Тренирует: буферизованные каналы.
// Сложность: easy
package main_test

import (
	"testing"
	"time"
)

func store(v int) int {
	ch := make(chan int)
	ch <- v
	return <-ch
}

func TestStore(t *testing.T) {
	done := make(chan int, 1)
	go func() { done <- store(7) }()
	select {
	case got := <-done:
		if got != 7 {
			t.Errorf("store = %d", got)
		}
	case <-time.After(time.Second):
		t.Fatal("store is blocked")
	}
}
