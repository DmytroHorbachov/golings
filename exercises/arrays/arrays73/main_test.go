// arrays73
// Make the tests pass!

// I AM NOT DONE
//
// Таблица кодов ошибок задана литералом [...] с явными индексами.
// Функция maxCode должна вернуть наибольший код, для которого есть сообщение.
// Тренирует: длина [...]-массива определяется наибольшим индексом + 1.
// Сложность: hard
package main_test

import "testing"

var messages = [...]string{
	1: "not found",
	4: "timeout",
	9: "internal",
}

func maxCode() int {
	return 3
}

func TestMaxCode(t *testing.T) {
	if got := maxCode(); got != 9 || messages[got] != "internal" {
		t.Errorf("maxCode = %d", got)
	}
}
