// concurrent_x090: Закрытие канала только для чтения
// Make the tests pass!
// I AM NOT DONE
//
// consumer получает канал только для чтения и пытается его закрыть.
// Код не компилируется.
// Тренирует: закрывать можно только канал, в который разрешена отправка.
// Сложность: hard
package main_test

import "testing"

func consumer(in <-chan int) int {
	s := 0
	for v := range in {
		s += v
	}
	close(in)
	return s
}

func producer(out chan<- int) {
	for i := 1; i <= 3; i++ {
		out <- i
	}
}

func TestConsumer(t *testing.T) {
	ch := make(chan int, 3)
	producer(ch)
	if got := consumer(ch); got != 6 {
		t.Errorf("sum = %d", got)
	}
}
