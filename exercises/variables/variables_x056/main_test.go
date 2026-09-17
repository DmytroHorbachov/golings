// variables_x056: Общее состояние счётчиков
// Make the tests pass!
// I AM NOT DONE
//
// Каждый счётчик, созданный newCounter, должен считать независимо.
// Сейчас все счётчики делят одно значение.
// Тренирует: время жизни переменных и захват в замыкании.
// Сложность: hard
package main_test

import "testing"

var count int

func newCounter() func() int {
	return func() int {
		count++
		return count
	}
}

func TestCounters(t *testing.T) {
	a := newCounter()
	b := newCounter()
	a()
	a()
	if got := a(); got != 3 {
		t.Errorf("a() third call = %d, want 3", got)
	}
	if got := b(); got != 1 {
		t.Errorf("b() first call = %d, want 1", got)
	}
}
