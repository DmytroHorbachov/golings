// concurrent24
// Make the tests pass!

// I AM NOT DONE
//
// startWorkers запускает воркеры, которые крутятся до отмены контекста,
// и возвращает функцию ожидания их завершения.
// Тренирует: ctx.Done() в цикле воркера.
// Сложность: medium
package main_test

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func startWorkers(ctx context.Context, n int, ticks *int64) func() {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				atomic.AddInt64(ticks, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	return wg.Wait
}

func TestStartWorkers(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	var ticks int64
	wait := startWorkers(ctx, 3, &ticks)
	time.Sleep(10 * time.Millisecond)
	cancel()
	done := make(chan struct{})
	go func() { wait(); close(done) }()
	select {
	case <-done:
		if atomic.LoadInt64(&ticks) == 0 {
			t.Errorf("workers did no work")
		}
	case <-time.After(2 * time.Second):
		t.Fatal("workers ignored cancellation")
	}
}
