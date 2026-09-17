// if_x074: Беззнаковая разность
// Make the tests pass!
// I AM NOT DONE
//
// shortage должна вернуть true, если запрошено больше, чем есть на складе.
// Условие never-true: разность беззнаковых чисел не бывает отрицательной.
// Тренирует: сравнение беззнаковых значений.
// Сложность: hard
package main_test

import "testing"

func shortage(stock, requested uint) bool {
	if stock-requested < 0 {
		return true
	}
	return false
}

func TestShortage(t *testing.T) {
	if !shortage(3, 5) {
		t.Errorf("shortage(3, 5) should be true")
	}
	if shortage(5, 5) || shortage(10, 1) {
		t.Errorf("no shortage expected")
	}
}
