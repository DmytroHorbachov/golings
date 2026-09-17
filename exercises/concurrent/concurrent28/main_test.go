// concurrent28
// Make the tests pass!

// I AM NOT DONE
//
// pending возвращает количество непрочитанных сообщений в буферизованном канале.
// Тренирует: len для каналов.
// Сложность: easy
package main_test

import "testing"

func pending(ch chan string) int {
	return cap(ch)
}

func TestPending(t *testing.T) {
	ch := make(chan string, 5)
	ch <- "a"
	ch <- "b"
	if pending(ch) != 2 {
		t.Errorf("pending = %d", pending(ch))
	}
}
