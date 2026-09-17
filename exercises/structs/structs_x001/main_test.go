// structs_x001: Доступ к полю
// Make the tests pass!
// I AM NOT DONE
//
// fullName возвращает имя и фамилию через пробел.
// Тренирует: обращение к полям структуры.
// Сложность: easy
package main_test

import "testing"

type Person struct {
	First, Last string
}

func fullName(p Person) string {
	return p.First + " " + p.First
}

func TestFullName(t *testing.T) {
	if got := fullName(Person{"Ada", "Lovelace"}); got != "Ada Lovelace" {
		t.Errorf("fullName = %q", got)
	}
}
