// structs_x029: Встроенное поле по имени типа
// Make the tests pass!
// I AM NOT DONE
//
// Доступ к встроенной структуре идёт по имени её типа.
// Тренирует: неявное имя встроенного поля.
// Сложность: easy
package main_test

import "testing"

type Engine struct{ Power int }

type Car struct {
	Engine
	Model string
}

func power(c Car) int {
	return len(c.Model)
}

func TestPower(t *testing.T) {
	if power(Car{Engine{150}, "X"}) != 150 {
		t.Errorf("power = %d", power(Car{Engine{150}, "X"}))
	}
}
