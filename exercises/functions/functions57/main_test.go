// functions57
// Make the tests pass!

// I AM NOT DONE
//
// safely должна перехватить панику и вернуть false.
// recover вызывается во вспомогательной функции, и паника не перехватывается.
// Тренирует: recover работает, только если вызван непосредственно отложенной функцией.
// Сложность: hard
package main_test

import "testing"

func tryRecover(ok *bool) {
	if r := recover(); r != nil {
		*ok = false
	}
}

func safely(f func()) (ok bool) {
	ok = true
	defer func() {
		tryRecover(&ok)
	}()
	f()
	return ok
}

func TestSafely(t *testing.T) {
	if !safely(func() {}) {
		t.Errorf("safely(no panic) = false, want true")
	}
	if safely(func() { panic("oops") }) {
		t.Errorf("safely(panic) = true, want false")
	}
}
