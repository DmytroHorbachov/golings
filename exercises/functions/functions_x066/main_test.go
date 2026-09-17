// functions_x066: Минимум с ошибкой
// Make the tests pass!
// I AM NOT DONE
//
// minOf должна вернуть минимум аргументов или ошибку, если аргументов нет.
// Тренирует: вариативные функции и возврат (значение, error).
// Сложность: medium
package main_test

import (
	"errors"
	"testing"
)

func minOf(nums ...int) (int, error) {
	m := nums[0]
	for _, n := range nums[1:] {
		if n < m {
			m = n
		}
	}
	return m, nil
}

func TestMinOf(t *testing.T) {
	_ = errors.New
	if m, err := minOf(4, -2, 9); err != nil || m != -2 {
		t.Errorf("minOf(4,-2,9) = %d, %v", m, err)
	}
	if _, err := minOf(); err == nil {
		t.Errorf("minOf() should return an error")
	}
}
