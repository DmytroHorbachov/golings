// concurrent38
// Make the tests pass!

// I AM NOT DONE
//
// sumUntilClosed читает из канала в цикле с select и складывает значения.
// После закрытия канала чтение сразу возвращает ноль, и цикл не заканчивается.
// Тренирует: признак ok при чтении в select.
// Сложность: hard
package main_test

import (
	"testing"
	"time"
)

func sumUntilClosed(ch <-chan int, abort <-chan struct{}) (int, bool) {
	sum := 0
	for {
		select {
		case v := <-ch:
			sum += v
		case <-abort:
			return sum, false
		}
	}
}

func TestSumUntilClosed(t *testing.T) {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	close(ch)
	abort := make(chan struct{})
	go func() {
		time.Sleep(200 * time.Millisecond)
		close(abort)
	}()
	sum, completed := sumUntilClosed(ch, abort)
	if sum != 3 || !completed {
		t.Errorf("sum = %d, completed = %v", sum, completed)
	}
}
