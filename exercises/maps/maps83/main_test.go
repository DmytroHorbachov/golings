// maps83
// Make the tests pass!

// I AM NOT DONE
//
// promote увеличивает уровень сотрудника, хранящегося в map.
// Код не компилируется: поле значения map нельзя изменить напрямую.
// Тренирует: значения map не адресуемы.
// Сложность: hard
package main_test

import "testing"

type Employee struct {
	Name  string
	Level int
}

func promote(staff map[string]Employee, name string) {
	staff[name].Level++
}

func TestPromote(t *testing.T) {
	staff := map[string]*Employee{"ann": {Name: "ann", Level: 1}}
	promote(staff, "ann")
	if staff["ann"].Level != 2 {
		t.Errorf("Level = %d, want 2", staff["ann"].Level)
	}
}
