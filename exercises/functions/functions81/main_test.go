// functions81
// Make the tests pass!

// I AM NOT DONE
//
// Функция statusOf должна достать код из ошибки типа *HTTPError в цепочке.
// Если такой ошибки нет, вернуть 0.
// Тренирует: errors.As и пользовательские типы ошибок.
// Сложность: medium
package main_test

import (
	"errors"
	"fmt"
	"testing"
)

type HTTPError struct{ Code int }

func (e *HTTPError) Error() string { return fmt.Sprintf("http %d", e.Code) }

func statusOf(err error) int {
	if he, ok := err.(*HTTPError); ok {
		return he.Code
	}
	return 0
}

func TestStatusOf(t *testing.T) {
	err := fmt.Errorf("fetch: %w", &HTTPError{Code: 404})
	if got := statusOf(err); got != 404 {
		t.Errorf("statusOf(wrapped 404) = %d, want 404", got)
	}
	if got := statusOf(errors.New("boom")); got != 0 {
		t.Errorf("statusOf(boom) = %d, want 0", got)
	}
}
