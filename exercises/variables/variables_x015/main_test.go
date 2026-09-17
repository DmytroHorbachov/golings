// variables_x015: Нулевое значение строки
// Make the tests pass!
// I AM NOT DONE
//
// Функция isEmpty должна сообщать, что строка пуста.
// Тренирует: нулевое значение строки — это "", а не nil.
// Сложность: easy
package main_test

import "testing"

func isEmpty(s string) bool {
	var empty string
	return s != empty
}

func TestIsEmpty(t *testing.T) {
	if !isEmpty("") {
		t.Errorf("isEmpty(\"\") = false, want true")
	}
	if isEmpty("go") {
		t.Errorf("isEmpty(\"go\") = true, want false")
	}
}
