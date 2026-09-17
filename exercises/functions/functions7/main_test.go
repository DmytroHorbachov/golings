// functions7
// Make the tests pass!

// I AM NOT DONE
//
// Функция product перемножает все аргументы; без аргументов результат 1.
// Сейчас произведение всегда 0.
// Тренирует: вариативные функции и начальное значение аккумулятора.
// Сложность: easy
package main_test

import "testing"

func product(nums ...int) int {
	result := 0
	for _, n := range nums {
		result *= n
	}
	return result
}

func TestProduct(t *testing.T) {
	if got := product(2, 3, 4); got != 24 {
		t.Errorf("product(2,3,4) = %d, want 24", got)
	}
	if got := product(); got != 1 {
		t.Errorf("product() = %d, want 1", got)
	}
}
