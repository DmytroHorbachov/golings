// functions_x033: Голый return
// Make the tests pass!
// I AM NOT DONE
//
// split делит сумму на две части: x = sum*4/9 и y = sum - x.
// Тренирует: именованные результаты и return без выражений.
// Сложность: easy
package main_test

import "testing"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum + x
	return
}

func TestSplit(t *testing.T) {
	x, y := split(18)
	if x != 8 || y != 10 {
		t.Errorf("split(18) = %d, %d; want 8, 10", x, y)
	}
}
