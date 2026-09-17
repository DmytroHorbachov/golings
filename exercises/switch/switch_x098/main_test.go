// switch_x098: Именованный тип в type switch
// Make the tests pass!
// I AM NOT DONE
//
// isInt должна считать целыми и int, и значения типа ID (type ID int).
// Тренирует: type switch сравнивает точный тип; ID — не int.
// Сложность: hard
package main_test

import "testing"

type ID int

func isInt(v interface{}) bool {
	switch v.(type) {
	case int:
		return true
	}
	return false
}

func TestIsInt(t *testing.T) {
	if !isInt(5) || !isInt(ID(7)) {
		t.Errorf("int and ID should be ints")
	}
	if isInt("5") {
		t.Errorf("string is not int")
	}
}
