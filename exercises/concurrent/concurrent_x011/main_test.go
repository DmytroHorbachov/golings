// concurrent_x011: Атомарный счётчик
// Make the tests pass!
// I AM NOT DONE
//
// countHits увеличивает счётчик из многих горутин без синхронизации.
// Тренирует: sync/atomic.
// Сложность: easy
package main_test

import (
	"sync"
	"sync/atomic"
	"testing"
)

func countHits(n int) int64 {
	var hits int64
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			hits++
		}()
	}
	wg.Wait()
	return atomic.LoadInt64(&hits)
}

func TestCountHits(t *testing.T) {
	if got := countHits(200); got != 200 {
		t.Errorf("hits = %d", got)
	}
}
