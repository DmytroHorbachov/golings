// concurrent44
// Make the tests pass!

// I AM NOT DONE
//
// Конфигурация публикуется через atomic.Value; Current возвращает её.
// Тренирует: Store и Load с утверждением типа.
// Сложность: easy
package main_test

import (
	"sync/atomic"
	"testing"
)

type Settings struct{ Level int }

var current atomic.Value

func Current() Settings {
	return Settings{}
}

func TestCurrent(t *testing.T) {
	current.Store(Settings{Level: 3})
	if Current().Level != 3 {
		t.Errorf("Current = %+v", Current())
	}
}
