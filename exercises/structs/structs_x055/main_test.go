// structs_x055: Повышение зарплаты
// Make the tests pass!
// I AM NOT DONE
//
// RaiseAll повышает зарплату всем сотрудникам отдела на процент.
// Тренирует: изменение элементов среза структур через индекс.
// Сложность: medium
package main_test

import "testing"

type Employee struct {
	Dept   string
	Salary int
}

func RaiseAll(staff []Employee, dept string, pct int) {
	for _, e := range staff {
		e.Salary += e.Salary * pct / 100
	}
}

func TestRaiseAll(t *testing.T) {
	staff := []Employee{{"it", 1000}, {"hr", 1000}}
	RaiseAll(staff, "it", 10)
	if staff[0].Salary != 1100 || staff[1].Salary != 1000 {
		t.Errorf("staff = %+v", staff)
	}
}
