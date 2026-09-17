// anonymous_functions_x089: sync.Once и паника
// Make the tests pass!
// I AM NOT DONE
//
// Конфигурация загружается один раз через sync.Once. Первая попытка паникует,
// и Once больше никогда не вызывает литерал — конфигурация остаётся пустой.
// Тренирует: Once считает вызов выполненным, даже если литерал запаниковал.
// Сложность: hard
package main_test

import (
	"sync"
	"testing"
)

type Loader struct {
	once  sync.Once
	value string
	load  func() string
}

func (l *Loader) Get() string {
	func() {
		defer func() { _ = recover() }()
		l.once.Do(func() {
			l.value = l.load()
		})
	}()
	return l.value
}

func TestLoader(t *testing.T) {
	l := &Loader{load: func() string { panic("no file") }}
	l.Get()
	if got := l.Get(); got != "default" {
		t.Errorf("Get = %q, want default", got)
	}
}
