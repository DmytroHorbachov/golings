// switch46
// Make the tests pass!

// I AM NOT DONE
//
// levels должна вернуть список ролей, которые включает данная роль:
// admin включает editor и viewer, editor включает viewer.
// Код не компилируется.
// Тренирует: fallthrough нельзя ставить в последней ветке switch.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

func levels(role string) []string {
	var out []string
	switch role {
	case "admin":
		out = append(out, "admin")
		fallthrough
	case "editor":
		out = append(out, "editor")
		fallthrough
	case "viewer":
		out = append(out, "viewer")
		fallthrough
	}
	return out
}

func TestLevels(t *testing.T) {
	if got := levels("admin"); !reflect.DeepEqual(got, []string{"admin", "editor", "viewer"}) {
		t.Errorf("levels(admin) = %v", got)
	}
	if got := levels("viewer"); !reflect.DeepEqual(got, []string{"viewer"}) {
		t.Errorf("levels(viewer) = %v", got)
	}
	if got := levels("guest"); len(got) != 0 {
		t.Errorf("levels(guest) = %v", got)
	}
}
