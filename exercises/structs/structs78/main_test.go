// structs78
// Make the tests pass!

// I AM NOT DONE
//
// Employee встраивает Person (с полем Name) и объявляет своё поле Name для
// должности. Функция должна вернуть имя человека.
// Тренирует: поле внешней структуры скрывает продвинутое поле.
// Сложность: hard
package main_test

import "testing"

type Person struct{ Name string }

type Employee struct {
	Person
	Name string
}

func personName(e Employee) string {
	return e.Name
}

func TestPersonName(t *testing.T) {
	e := Employee{Person{"Ann"}, "Engineer"}
	if personName(e) != "Ann" {
		t.Errorf("personName = %q", personName(e))
	}
}
