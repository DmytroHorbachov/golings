// if32
// Make the tests pass!

// I AM NOT DONE
//
// absInt64 возвращает модуль числа или ошибку, если модуль не представим.
// Для math.MinInt64 функция возвращает отрицательное число без ошибки.
// Тренирует: -MinInt64 переполняется и снова равен MinInt64.
// Сложность: hard
package main_test

import (
	"errors"
	"math"
	"testing"
)

var ErrRange = errors.New("out of range")

func absInt64(x int64) (int64, error) {
	if x < 0 {
		return -x, nil
	}
	return x, nil
}

func TestAbsInt64(t *testing.T) {
	if v, err := absInt64(-5); err != nil || v != 5 {
		t.Errorf("absInt64(-5) = %d, %v", v, err)
	}
	if v, err := absInt64(math.MaxInt64); err != nil || v != math.MaxInt64 {
		t.Errorf("absInt64(max) = %d, %v", v, err)
	}
	if _, err := absInt64(math.MinInt64); !errors.Is(err, ErrRange) {
		t.Errorf("absInt64(min) error = %v, want ErrRange", err)
	}
}
