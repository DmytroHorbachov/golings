// functions48
// Make the tests pass!

// I AM NOT DONE
//
// once(f) должна вернуть функцию, которая вызывает f только при первом вызове,
// а затем возвращает сохранённый результат.
// Тренирует: замыкания с флагом и кэшированным значением.
// Сложность: medium
package main_test

import "testing"

func once(f func() int) func() int {
	return func() int {
		return f()
	}
}

func TestOnce(t *testing.T) {
	calls := 0
	g := once(func() int { calls++; return 42 })
	if g() != 42 || g() != 42 || g() != 42 {
		t.Errorf("once() should always return 42")
	}
	if calls != 1 {
		t.Errorf("wrapped function called %d times, want 1", calls)
	}
}
