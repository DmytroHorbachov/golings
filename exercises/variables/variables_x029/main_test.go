// variables_x029: Беззнаковое вычитание
// Make the tests pass!
// I AM NOT DONE
//
// Функция должна вернуть, сколько товара осталось на складе, но не меньше нуля.
// Когда заказ больше остатка, получается огромное число.
// Тренирует: переполнение беззнаковых типов при вычитании.
// Сложность: hard
package main_test

import "testing"

func remaining(stock, order uint) uint {
	left := stock - order
	if left < 0 {
		return 0
	}
	return left
}

func TestRemaining(t *testing.T) {
	if got := remaining(10, 3); got != 7 {
		t.Errorf("remaining(10, 3) = %d, want 7", got)
	}
	if got := remaining(3, 10); got != 0 {
		t.Errorf("remaining(3, 10) = %d, want 0", got)
	}
}
