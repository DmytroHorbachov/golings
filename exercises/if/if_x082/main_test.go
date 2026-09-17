// if_x082: Приоритет && и
// Make the tests pass!
// I AM NOT DONE
//
// canEdit: редактировать могут администратор или владелец, но только
// если документ не заблокирован.
// Тренирует: && имеет более высокий приоритет, чем ||.
// Сложность: hard
package main_test

import "testing"

func canEdit(admin, owner, locked bool) bool {
	if admin || owner && !locked {
		return true
	}
	return false
}

func TestCanEdit(t *testing.T) {
	cases := []struct {
		admin, owner, locked, want bool
	}{
		{true, false, false, true}, {false, true, false, true},
		{true, false, true, false}, {false, true, true, false}, {false, false, false, false},
	}
	for _, c := range cases {
		if got := canEdit(c.admin, c.owner, c.locked); got != c.want {
			t.Errorf("canEdit(%v, %v, %v) = %v, want %v", c.admin, c.owner, c.locked, got, c.want)
		}
	}
}
