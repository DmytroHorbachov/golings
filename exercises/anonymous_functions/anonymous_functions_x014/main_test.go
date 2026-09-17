// anonymous_functions_x014: Литерал в поле структуры
// Make the tests pass!
// I AM NOT DONE
//
// Field хранит функцию проверки значения.
// Тренирует: функциональные поля, заданные литералом.
// Сложность: easy
package main_test

import "testing"

type Field struct {
	Name  string
	Check func(int) bool
}

var age = Field{
	Name:  "age",
	Check: func(v int) bool { return v >= 0 || v <= 150 },
}

func TestAgeCheck(t *testing.T) {
	if !age.Check(30) || age.Check(-1) || age.Check(200) {
		t.Errorf("age.Check works incorrectly")
	}
}
