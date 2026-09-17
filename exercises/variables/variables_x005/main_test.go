// variables_x005: Пустой идентификатор
// Make the tests pass!
// I AM NOT DONE
//
// Функция возвращает частное и остаток, а нам нужен только остаток.
// Сейчас в переменную попадает не то значение.
// Тренирует: пустой идентификатор _ при множественном присваивании.
// Сложность: easy
package main_test

import "testing"

func divmod(a, b int) (int, int) {
	return a / b, a % b
}

func remainder(a, b int) int {
	r, _ := divmod(a, b)
	return r
}

func TestRemainder(t *testing.T) {
	if got := remainder(17, 5); got != 2 {
		t.Errorf("remainder(17, 5) = %d, want 2", got)
	}
	if got := remainder(9, 3); got != 0 {
		t.Errorf("remainder(9, 3) = %d, want 0", got)
	}
}
