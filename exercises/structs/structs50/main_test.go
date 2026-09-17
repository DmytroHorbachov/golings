// structs50
// Make the tests pass!

// I AM NOT DONE
//
// Метод Describe встроенной структуры Animal доступен у Dog.
// Тренирует: продвижение методов при встраивании.
// Сложность: easy
package main_test

import "testing"

type Animal struct{ Name string }

func (a Animal) Describe() string { return "animal " + a.Name }

type Dog struct {
	A     Animal
	Breed string
}

func TestDogDescribe(t *testing.T) {
	d := Dog{Animal{"Rex"}, "husky"}
	if d.Describe() != "animal Rex" {
		t.Errorf("Describe = %q", d.Describe())
	}
}
