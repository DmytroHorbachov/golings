// functions29
// Make the tests pass!

// I AM NOT DONE
//
// counted(f) должна вернуть обёртку над f и функцию, сообщающую число вызовов.
// Тренирует: две функции, разделяющие одно состояние.
// Сложность: medium
package main_test

import "testing"

func counted(f func(int) int) (func(int) int, func() int) {
	calls := 0
	wrapped := func(x int) int {
		return f(x)
	}
	count := func() int { return 0 }
	return wrapped, count
}

func TestCounted(t *testing.T) {
	neg, count := counted(func(x int) int { return -x })
	neg(1)
	neg(2)
	if got := neg(3); got != -3 {
		t.Errorf("neg(3) = %d, want -3", got)
	}
	if got := count(); got != 3 {
		t.Errorf("count() = %d, want 3", got)
	}
}
