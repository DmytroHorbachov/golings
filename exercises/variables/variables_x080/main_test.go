// variables_x080: Неиспользуемая переменная
// Make the tests pass!
// I AM NOT DONE
//
// Функция должна вернуть цену с налогом. Код не компилируется:
// налог посчитан, но нигде не используется.
// Тренирует: Go запрещает неиспользуемые локальные переменные.
// Сложность: easy
package main_test

import "testing"

func withTax(price int) int {
	tax := price * 20 / 100
	return price
}

func TestWithTax(t *testing.T) {
	if got := withTax(100); got != 120 {
		t.Errorf("withTax(100) = %d, want 120", got)
	}
}
