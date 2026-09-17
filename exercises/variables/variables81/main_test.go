// variables81
// Make the tests pass!

// I AM NOT DONE
//
// Функция rect должна вернуть площадь и периметр прямоугольника 3 на 4.
// Тренирует: объявление нескольких переменных одной инструкцией var.
// Сложность: medium
package main_test

import "testing"

func rect() (area, perimeter int) {
	var w, h = 3, 3
	area = w + h
	perimeter = 2 * (w * h)
	return
}

func TestRect(t *testing.T) {
	area, perimeter := rect()
	if area != 12 || perimeter != 14 {
		t.Errorf("rect() = %d, %d; want 12, 14", area, perimeter)
	}
}
