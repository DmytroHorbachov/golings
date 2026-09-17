// switch_x081: type switch и обёрнутые ошибки
// Make the tests pass!
// I AM NOT DONE
//
// httpCode возвращает код из *StatusError в цепочке ошибок или 500.
// type switch не видит ошибку, обёрнутую через %w.
// Тренирует: type switch проверяет только верхний уровень ошибки.
// Сложность: hard
package main_test

import (
	"errors"
	"fmt"
	"testing"
)

type StatusError struct{ Code int }

func (e *StatusError) Error() string { return fmt.Sprintf("status %d", e.Code) }

func httpCode(err error) int {
	switch e := err.(type) {
	case *StatusError:
		return e.Code
	}
	return 500
}

func TestHTTPCode(t *testing.T) {
	_ = errors.As
	if got := httpCode(&StatusError{404}); got != 404 {
		t.Errorf("httpCode(404) = %d", got)
	}
	if got := httpCode(fmt.Errorf("handler: %w", &StatusError{403})); got != 403 {
		t.Errorf("httpCode(wrapped 403) = %d, want 403", got)
	}
	if got := httpCode(errors.New("boom")); got != 500 {
		t.Errorf("httpCode(boom) = %d", got)
	}
}
