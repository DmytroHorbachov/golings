// functions35
// Make the tests pass!

// I AM NOT DONE
//
// safeCall должна выполнить f и превратить панику в ошибку.
// Сейчас паника пролетает насквозь.
// Тренирует: defer, recover и именованный результат.
// Сложность: medium
package main_test

import (
	"fmt"
	"testing"
)

func safeCall(f func()) error {
	f()
	return nil
}

func TestSafeCall(t *testing.T) {
	_ = fmt.Sprint
	if err := safeCall(func() {}); err != nil {
		t.Errorf("safeCall(ok) = %v, want nil", err)
	}
	err := safeCall(func() { panic("boom") })
	if err == nil || err.Error() != "panic: boom" {
		t.Errorf("safeCall(panic) = %v, want panic: boom", err)
	}
}
