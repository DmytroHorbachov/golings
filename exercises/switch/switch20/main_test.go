// switch20
// Make the tests pass!

// I AM NOT DONE
//
// isOne должна вернуть true, если значение равно единице любого целого типа
// из списка int, int64, uint8.
// Тренирует: case 1 при switch по interface{} сравнивает и тип (int), и значение.
// Сложность: hard
package main_test

import "testing"

func isOne(v interface{}) bool {
	switch v {
	case 1:
		return true
	}
	return false
}

func TestIsOne(t *testing.T) {
	for _, v := range []interface{}{1, int64(1), uint8(1)} {
		if !isOne(v) {
			t.Errorf("isOne(%T(1)) = false", v)
		}
	}
	for _, v := range []interface{}{2, "1", 1.0} {
		if isOne(v) {
			t.Errorf("isOne(%T(%v)) = true", v, v)
		}
	}
}
