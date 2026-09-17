// concurrent11
// Make the tests pass!

// I AM NOT DONE
//
// limited запускает задачи так, чтобы одновременно работали не больше 2.
// Размер семафора задан неверно.
// Тренирует: буферизованный канал как семафор.
// Сложность: easy
package main_test

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func limited(n int) int32 {
	sem := make(chan struct{}, n)
	var cur, peak int32
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			c := atomic.AddInt32(&cur, 1)
			for {
				p := atomic.LoadInt32(&peak)
				if c <= p || atomic.CompareAndSwapInt32(&peak, p, c) {
					break
				}
			}
			time.Sleep(20 * time.Millisecond)
			atomic.AddInt32(&cur, -1)
		}()
	}
	wg.Wait()
	return peak
}

func TestLimited(t *testing.T) {
	if peak := limited(6); peak != 2 {
		t.Errorf("peak concurrency = %d, want 2", peak)
	}
}
