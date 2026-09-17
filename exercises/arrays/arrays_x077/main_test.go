// arrays_x077: Копия с общим срезом
// Make the tests pass!
// I AM NOT DONE
//
// cloneTeam копирует команду (массив игроков), чтобы изменения в копии
// не затрагивали оригинал. Однако навыки игроков оказываются общими.
// Тренирует: копирование массива структур — поверхностное.
// Сложность: hard
package main_test

import "testing"

type Player struct {
	Name   string
	Skills []string
}

func cloneTeam(t [2]Player) [2]Player {
	c := t
	return c
}

func TestCloneTeam(t *testing.T) {
	team := [2]Player{{"ann", []string{"pass"}}, {"bob", []string{"shoot"}}}
	c := cloneTeam(team)
	c[0].Skills[0] = "dribble"
	c[1].Name = "rob"
	if team[0].Skills[0] != "pass" || team[1].Name != "bob" {
		t.Errorf("original changed: %v", team)
	}
}
