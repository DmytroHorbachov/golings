// switch_x082: switch err и обёртки
// Make the tests pass!
// I AM NOT DONE
//
// retryable возвращает true для ErrTimeout и ErrBusy, даже если они обёрнуты.
// switch по значению ошибки обёртки не видит.
// Тренирует: сравнение ошибок в case выполняется через ==.
// Сложность: hard
package main_test

import (
	"errors"
	"fmt"
	"testing"
)

var (
	ErrTimeout = errors.New("timeout")
	ErrBusy    = errors.New("busy")
)

func retryable(err error) bool {
	switch err {
	case ErrTimeout, ErrBusy:
		return true
	}
	return false
}

func TestRetryable(t *testing.T) {
	if !retryable(ErrBusy) || !retryable(fmt.Errorf("db: %w", ErrTimeout)) {
		t.Errorf("timeout and busy errors are retryable")
	}
	if retryable(errors.New("timeout")) || retryable(nil) {
		t.Errorf("other errors are not retryable")
	}
}
