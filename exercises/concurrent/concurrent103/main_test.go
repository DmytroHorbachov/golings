// concurrent103
// Make the tests pass!

// I AM NOT DONE
//
// Главная горутина ждёт воркеров, а затем отправляет им задачу, которую
// они ждут. Возникает взаимоблокировка.
// Тренирует: порядок операций при синхронизации.
// Сложность: hard
package main_test

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func startAndFeed(workers int) int32 {
	tasks := make(chan int)
	var sum int32
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&sum, int32(<-tasks))
		}()
	}
	wg.Wait()
	for i := 0; i < workers; i++ {
		tasks <- 2
	}
	return atomic.LoadInt32(&sum)
}

func TestStartAndFeed(t *testing.T) {
	res := make(chan int32, 1)
	go func() { res <- startAndFeed(3) }()
	select {
	case got := <-res:
		if got != 6 {
			t.Errorf("sum = %d", got)
		}
	case <-time.After(time.Second):
		t.Fatal("deadlock: tasks are sent after waiting")
	}
}
