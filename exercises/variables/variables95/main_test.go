// variables95
// Make the tests pass!

// I AM NOT DONE
//
// Функция должна увеличить значение и вернуть его. Код не компилируется.
// Тренирует: := требует хотя бы одну новую переменную слева.
// Сложность: easy
package main_test

import "testing"

func addTen(start int) int {
	total := start
	total := total + 10
	return total
}

func TestAddTen(t *testing.T) {
	if got := addTen(5); got != 15 {
		t.Errorf("addTen(5) = %d, want 15", got)
	}
}
