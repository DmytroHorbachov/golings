// concurrent_x081: Атомарность наполовину
// Make the tests pass!
// I AM NOT DONE
//
// Счётчик увеличивается через atomic, но читается напрямую.
// Детектор гонок сообщает об ошибке.
// Тренирует: все обращения к атомарной переменной — только через atomic.
// Сложность: hard
package main_test

import (
	"sync"
	"sync/atomic"
	"testing"
)

type Meter struct{ bytes int64 }

func (m *Meter) Add(n int64) { atomic.AddInt64(&m.bytes, n) }

func (m *Meter) Total() int64 {
	return m.bytes
}

func TestMeter(t *testing.T) {
	var m Meter
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(2)
		go func() { defer wg.Done(); m.Add(10) }()
		go func() { defer wg.Done(); _ = m.Total() }()
	}
	wg.Wait()
	if m.Total() != 500 {
		t.Errorf("Total = %d", m.Total())
	}
}
