// concurrent23
// Make the tests pass!

// I AM NOT DONE
//
// Конфигурация хранится в atomic.Value; при обновлении сохраняется значение
// другого типа, и Store паникует.
// Тренирует: atomic.Value требует одного конкретного типа для всех Store.
// Сложность: hard
package main_test

import (
	"sync/atomic"
	"testing"
)

type Config struct{ Level int }

type Holder struct{ v atomic.Value }

func (h *Holder) Set(level int) {
	if level == 0 {
		h.v.Store(Config{})
		return
	}
	h.v.Store(&Config{Level: level})
}

func (h *Holder) Level() int { return h.v.Load().(*Config).Level }

func TestHolder(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("panic: %v", r)
		}
	}()
	var h Holder
	h.Set(3)
	h.Set(0)
	if h.Level() != 0 {
		t.Errorf("Level = %d", h.Level())
	}
}
