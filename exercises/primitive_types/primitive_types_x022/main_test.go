// primitive_types_x022: bool в число
// Make the tests pass!
// I AM NOT DONE
//
// boolToInt должна вернуть 1 для true и 0 для false.
// Код не компилируется: bool нельзя преобразовать в int.
// Тренирует: отсутствие неявных преобразований bool.
// Сложность: medium
package main_test

import "testing"

func boolToInt(b bool) int {
	return int(b)
}

func TestBoolToInt(t *testing.T) {
	if boolToInt(true) != 1 || boolToInt(false) != 0 {
		t.Errorf("boolToInt is wrong")
	}
}
