// variables75
// Make the tests pass!

// I AM NOT DONE
//
// Функция almostEqual должна считать 0.1+0.2 и 0.3 равными.
// Точное сравнение чисел с плавающей точкой здесь не работает.
// Тренирует: погрешность вычислений с float64 и сравнение с эпсилоном.
// Сложность: medium
package main_test

import (
	"math"
	"testing"
)

const epsilon = 1e-9

func almostEqual(a, b float64) bool {
	_ = math.Abs
	return a == b
}

func TestAlmostEqual(t *testing.T) {
	x, y := 0.1, 0.2
	if !almostEqual(x+y, 0.3) {
		t.Errorf("0.1+0.2 should be almost equal to 0.3")
	}
	if almostEqual(0.3, 0.31) {
		t.Errorf("0.3 and 0.31 should not be almost equal")
	}
}
