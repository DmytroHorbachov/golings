// arrays_x046: Строка судоку
// Make the tests pass!
// I AM NOT DONE
//
// validRow проверяет, что в строке судоку цифры 1–9 не повторяются (0 — пустая клетка).
// Тренирует: массив-счётчик [10]bool.
// Сложность: medium
package main_test

import "testing"

func validRow(row [9]int) bool {
	var seen [10]bool
	for _, v := range row {
		if seen[v] {
			return false
		}
	}
	return true
}

func TestValidRow(t *testing.T) {
	if !validRow([9]int{5, 3, 0, 0, 7, 0, 0, 0, 0}) {
		t.Errorf("row with blanks should be valid")
	}
	if validRow([9]int{5, 3, 5, 0, 7, 0, 0, 0, 0}) {
		t.Errorf("row with duplicate 5 should be invalid")
	}
}
