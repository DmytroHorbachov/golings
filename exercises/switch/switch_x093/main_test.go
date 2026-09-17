// switch_x093: Вычисленное дробное значение
// Make the tests pass!
// I AM NOT DONE
//
// coinName определяет монету по номиналу, вычисленному как сумма долей.
// Для 0.1+0.2 ветка 0.3 не срабатывает.
// Тренирует: switch сравнивает float точно, а вычисления дают погрешность.
// Сложность: hard
package main_test

import (
	"math"
	"testing"
)

func coinName(value float64) string {
	switch value {
	case 0.3:
		return "thirty"
	case 0.5:
		return "fifty"
	}
	return "unknown"
}

func TestCoinName(t *testing.T) {
	_ = math.Round
	a, b := 0.1, 0.2
	if got := coinName(a + b); got != "thirty" {
		t.Errorf("coinName(0.1+0.2) = %s, want thirty", got)
	}
	if got := coinName(0.25 + 0.25); got != "fifty" {
		t.Errorf("coinName(0.5) = %s", got)
	}
	if got := coinName(0.7); got != "unknown" {
		t.Errorf("coinName(0.7) = %s", got)
	}
}
