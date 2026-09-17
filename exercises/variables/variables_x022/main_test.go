// variables_x022: Область видимости if
// Make the tests pass!
// I AM NOT DONE
//
// Функция должна вернуть скидку: 10 для суммы от 1000, иначе 0.
// Переменная объявлена внутри блока if и снаружи недоступна.
// Тренирует: области видимости переменных.
// Сложность: medium
package main_test

import "testing"

func discount(total int) int {
	if total >= 1000 {
		d := 10
	}
	return d
}

func TestDiscount(t *testing.T) {
	if got := discount(1500); got != 10 {
		t.Errorf("discount(1500) = %d, want 10", got)
	}
	if got := discount(999); got != 0 {
		t.Errorf("discount(999) = %d, want 0", got)
	}
}
