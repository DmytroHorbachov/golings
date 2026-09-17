// functions40
// Make the tests pass!

// I AM NOT DONE
//
// isTimeout должна распознать ErrTimeout, даже если ошибка обёрнута.
// Тренирует: errors.Is вместо сравнения через ==.
// Сложность: easy
package main_test

import (
	"errors"
	"fmt"
	"testing"
)

var ErrTimeout = errors.New("timeout")

func isTimeout(err error) bool {
	return err == ErrTimeout
}

func TestIsTimeout(t *testing.T) {
	wrapped := fmt.Errorf("call api: %w", ErrTimeout)
	if !isTimeout(ErrTimeout) || !isTimeout(wrapped) {
		t.Errorf("isTimeout should detect direct and wrapped ErrTimeout")
	}
	if isTimeout(errors.New("timeout")) {
		t.Errorf("isTimeout should not match a different error with the same text")
	}
}
