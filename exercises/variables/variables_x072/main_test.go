// variables_x072: Комплексные числа
// Make the tests pass!
// I AM NOT DONE
//
// Функция должна вернуть модуль комплексного числа 3+4i и его действительную часть.
// Тренирует: тип complex128 и функции real, imag, cmplx.Abs.
// Сложность: medium
package main_test

import (
	"math/cmplx"
	"testing"
)

func describe() (float64, float64) {
	z := complex(4, 4)
	return cmplx.Abs(z), imag(z)
}

func TestDescribe(t *testing.T) {
	abs, re := describe()
	if abs != 5 {
		t.Errorf("abs = %v, want 5", abs)
	}
	if re != 3 {
		t.Errorf("real part = %v, want 3", re)
	}
}
