// switch_x032: Тип nil
// Make the tests pass!
// I AM NOT DONE
//
// typeName возвращает "nil" для nil-интерфейса, "bool" для bool, иначе "other".
// Тренирует: case nil в type switch.
// Сложность: easy
package main_test

import "testing"

func typeName(v interface{}) string {
	switch v.(type) {
	case error:
		return "nil"
	case bool:
		return "bool"
	}
	return "other"
}

func TestTypeName(t *testing.T) {
	if got := typeName(nil); got != "nil" {
		t.Errorf("typeName(nil) = %s", got)
	}
	if got := typeName(true); got != "bool" {
		t.Errorf("typeName(true) = %s", got)
	}
	if got := typeName(3); got != "other" {
		t.Errorf("typeName(3) = %s", got)
	}
}
