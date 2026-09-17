// concurrent_x059: Замер максимальной параллельности
// Make the tests pass!
// I AM NOT DONE
//
// tracker отмечает вход и выход задач и запоминает пиковое число одновременных.
// Тренирует: мьютекс для нескольких связанных полей.
// Сложность: medium
package main_test

import (
	"sync"
	"testing"
	"time"
)

type tracker struct {
	mu        sync.Mutex
	cur, peak int
}

func (t *tracker) Enter() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.peak++
}

func (t *tracker) Leave() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.cur--
}

func TestTracker(t *testing.T) {
	var tr tracker
	sem := make(chan struct{}, 3)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sem <- struct{}{}
			tr.Enter()
			time.Sleep(10 * time.Millisecond)
			tr.Leave()
			<-sem
		}()
	}
	wg.Wait()
	if tr.peak != 3 || tr.cur != 0 {
		t.Errorf("peak = %d, cur = %d", tr.peak, tr.cur)
	}
}
