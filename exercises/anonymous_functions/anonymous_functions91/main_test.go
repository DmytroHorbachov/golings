// anonymous_functions91
// Make the tests pass!

// I AM NOT DONE
//
// safeGo запускает функцию в горутине и должна превратить её панику в ошибку.
// recover в вызывающей функции панику из другой горутины не ловит.
// Тренирует: recover работает только в той горутине, где произошла паника.
// Сложность: hard
package main_test

import (
	"fmt"
	"testing"
)

func safeGo(f func()) (err error) {
	done := make(chan error, 1)
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	go func() {
		f()
		done <- nil
	}()
	return <-done
}

func TestSafeGo(t *testing.T) {
	if err := safeGo(func() {}); err != nil {
		t.Errorf("safeGo(ok) = %v", err)
	}
	if err := safeGo(func() { panic("boom") }); err == nil || err.Error() != "panic: boom" {
		t.Errorf("safeGo(panic) = %v", err)
	}
}
