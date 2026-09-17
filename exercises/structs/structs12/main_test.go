// structs12
// Make the tests pass!

// I AM NOT DONE
//
// IsAdult сообщает, что человеку есть 18.
// Тренирует: методы, возвращающие bool.
// Сложность: easy
package main_test

import "testing"

type Person struct{ Age int }

func (p Person) IsAdult() bool {
	return p.Age > 18
}

func TestIsAdult(t *testing.T) {
	if !(Person{18}).IsAdult() || (Person{17}).IsAdult() {
		t.Errorf("IsAdult works incorrectly")
	}
}
