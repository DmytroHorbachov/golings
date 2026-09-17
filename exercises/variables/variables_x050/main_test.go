// variables_x050: Точность float32
// Make the tests pass!
// I AM NOT DONE
//
// Функция хранит идентификаторы как float32 и сравнивает их.
// Разные большие идентификаторы вдруг оказываются равными.
// Тренирует: ограниченная точность float32 для целых чисел.
// Сложность: hard
package main_test

import "testing"

func sameID(a, b int64) bool {
	x, y := float32(a), float32(b)
	return x == y
}

func TestSameID(t *testing.T) {
	if sameID(16777216, 16777217) {
		t.Errorf("16777216 and 16777217 are different ids")
	}
	if !sameID(42, 42) {
		t.Errorf("42 and 42 are the same id")
	}
}
