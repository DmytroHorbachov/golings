// variables65
// Make the tests pass!

// I AM NOT DONE
//
// Функция hypot должна вернуть длину гипотенузы для целых катетов.
// Код не компилируется: math.Sqrt принимает float64.
// Тренирует: явные преобразования числовых типов.
// Сложность: easy
package main_test

import (
	"math"
	"testing"
)

func hypot(a, b int) float64 {
	sq := a*a + b*b
	return math.Sqrt(sq)
}

func TestHypot(t *testing.T) {
	if got := hypot(3, 4); got != 5 {
		t.Errorf("hypot(3, 4) = %v, want 5", got)
	}
}
