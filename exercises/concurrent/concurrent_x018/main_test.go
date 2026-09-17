// concurrent_x018: Ёмкость канала
// Make the tests pass!
// I AM NOT DONE
//
// freeSlots возвращает, сколько ещё сообщений поместится в буфер.
// Тренирует: cap и len для каналов.
// Сложность: easy
package main_test

import "testing"

func freeSlots(ch chan int) int {
	return len(ch) - cap(ch)
}

func TestFreeSlots(t *testing.T) {
	ch := make(chan int, 4)
	ch <- 1
	if freeSlots(ch) != 3 {
		t.Errorf("freeSlots = %d", freeSlots(ch))
	}
}
