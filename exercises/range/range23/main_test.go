// range23
// Make the tests pass!

// I AM NOT DONE
//
// consume читает сообщения до получения "quit". break внутри select
// не выходит из цикла, и функция зависает.
// Тренирует: break относится к ближайшему for, switch или select.
// Сложность: hard
package main_test

import (
	"testing"
	"time"
)

func consume(msgs <-chan string) int {
	n := 0
	for {
		select {
		case m := <-msgs:
			if m == "quit" {
				break
			}
			n++
		}
	}
	return n
}

func TestConsume(t *testing.T) {
	msgs := make(chan string, 4)
	msgs <- "a"
	msgs <- "b"
	msgs <- "quit"
	done := make(chan int, 1)
	go func() { done <- consume(msgs) }()
	select {
	case n := <-done:
		if n != 2 {
			t.Errorf("consume = %d, want 2", n)
		}
	case <-time.After(time.Second):
		t.Fatal("consume did not stop on quit")
	}
}
