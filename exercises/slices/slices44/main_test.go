// slices44
// Make the tests pass!

// I AM NOT DONE
//
// columnSum суммирует столбец c в «рваной» таблице, где строки разной длины.
// Для коротких строк функция паникует.
// Тренирует: у [][]T каждая строка имеет свою длину.
// Сложность: hard
package main_test

import "testing"

func columnSum(rows [][]int, c int) int {
	sum := 0
	for _, r := range rows {
		if c < len(rows[0]) {
			sum += r[c]
		}
	}
	return sum
}

func TestColumnSum(t *testing.T) {
	rows := [][]int{{1, 2, 3}, {4}, {5, 6}}
	if got := columnSum(rows, 1); got != 8 {
		t.Errorf("columnSum(1) = %d, want 8", got)
	}
	if got := columnSum(rows, 0); got != 10 {
		t.Errorf("columnSum(0) = %d, want 10", got)
	}
}
