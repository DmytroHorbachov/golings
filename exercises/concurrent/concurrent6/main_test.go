// concurrent6
// Make the tests pass!

// I AM NOT DONE
//
// drain должна прекратить обработку сразу, если сигнал quit уже получен,
// даже когда в очереди есть работа. Обычный select выбирает ветку случайно.
// Тренирует: select не гарантирует порядок; приоритет требует отдельной проверки.
// Сложность: hard
package main_test

import "testing"

func drain(work <-chan int, quit <-chan struct{}) int {
	n := 0
	for {
		select {
		case <-quit:
			return n
		case <-work:
			n++
		}
	}
}

func TestDrainPriority(t *testing.T) {
	for trial := 0; trial < 20; trial++ {
		work := make(chan int, 100)
		for i := 0; i < 100; i++ {
			work <- i
		}
		quit := make(chan struct{})
		close(quit)
		if n := drain(work, quit); n != 0 {
			t.Fatalf("trial %d: processed %d items after quit", trial, n)
		}
	}
}
