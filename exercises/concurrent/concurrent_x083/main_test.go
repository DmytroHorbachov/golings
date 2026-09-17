// concurrent_x083: Add внутри горутины
// Make the tests pass!
// I AM NOT DONE
//
// wg.Add вызывается внутри запущенной горутины, и Wait возвращается раньше,
// чем работа начата.
// Тренирует: Add нужно вызывать до старта горутины.
// Сложность: hard
package main_test

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func runJobs(n int, gate <-chan struct{}) int32 {
	var done int32
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		go func() {
			<-gate
			wg.Add(1)
			defer wg.Done()
			atomic.AddInt32(&done, 1)
		}()
	}
	finished := make(chan struct{})
	go func() {
		wg.Wait()
		close(finished)
	}()
	<-finished
	return atomic.LoadInt32(&done)
}

func TestRunJobs(t *testing.T) {
	gate := make(chan struct{})
	res := make(chan int32, 1)
	go func() { res <- runJobs(5, gate) }()
	time.Sleep(50 * time.Millisecond)
	close(gate)
	if got := <-res; got != 5 {
		t.Errorf("done = %d, want 5", got)
	}
}
