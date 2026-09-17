// anonymous_functions_x095: Аргументы go вычисляются сразу
// Make the tests pass!
// I AM NOT DONE
//
// startJob запускает обработку в фоне, но функция загрузки аргумента
// выполняется в текущей горутине и блокирует вызывающего.
// Тренирует: в go f(x()) вызов x() происходит до старта горутины.
// Сложность: hard
package main_test

import (
	"testing"
	"time"
)

func startJob(load func() int, results chan<- int) {
	process := func(v int) { results <- v * 2 }
	go process(load())
}

func TestStartJob(t *testing.T) {
	release := make(chan struct{})
	results := make(chan int, 1)
	load := func() int { <-release; return 21 }
	started := make(chan struct{})
	go func() {
		startJob(load, results)
		close(started)
	}()
	select {
	case <-started:
	case <-time.After(time.Second):
		close(release)
		t.Fatal("startJob blocked the caller")
	}
	close(release)
	if got := <-results; got != 42 {
		t.Errorf("result = %d", got)
	}
}
