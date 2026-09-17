// switch_x070: fallthrough в type switch
// Make the tests pass!
// I AM NOT DONE
//
// numeric должна вернуть true для int и int64 значений.
// Код не компилируется: fallthrough в type switch запрещён.
// Тренирует: ограничения type switch.
// Сложность: hard
package main_test

import "testing"

func numeric(v interface{}) bool {
	switch v.(type) {
	case int:
		fallthrough
	case int64:
		return true
	}
	return false
}

func TestNumeric(t *testing.T) {
	if !numeric(1) || !numeric(int64(2)) {
		t.Errorf("int and int64 are numeric")
	}
	if numeric("3") {
		t.Errorf("string is not numeric")
	}
}
