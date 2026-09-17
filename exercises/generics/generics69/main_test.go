// generics69
// Make the tests pass!

// I AM NOT DONE
//
// Abs возвращает модуль числа любого знакового типа.
// Тренирует: ограничение на знаковые числа.
// Сложность: easy
package main_test

import "testing"

type Signed interface {
	~int | ~int8 | ~int32 | ~int64 | ~float64
}

func Abs[T Signed](v T) T {
	if v > 0 {
		return -v
	}
	return v
}

func TestAbs(t *testing.T) {
	if Abs(-3) != 3 || Abs(int8(4)) != 4 || Abs(-2.5) != 2.5 {
		t.Errorf("Abs works incorrectly")
	}
}
