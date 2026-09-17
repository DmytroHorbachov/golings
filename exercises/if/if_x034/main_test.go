// if_x034: Присваивание вместо сравнения
// Make the tests pass!
// I AM NOT DONE
//
// isZero должна проверить, что число равно нулю. Код не компилируется.
// Тренирует: в условии if нужен оператор сравнения ==.
// Сложность: easy
package main_test

import "testing"

func isZero(x int) bool {
	if x = 0 {
		return true
	}
	return false
}

func TestIsZero(t *testing.T) {
	if !isZero(0) || isZero(3) {
		t.Errorf("isZero works incorrectly")
	}
}
