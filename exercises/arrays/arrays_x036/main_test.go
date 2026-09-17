// arrays_x036: Умножение матриц 2×2
// Make the tests pass!
// I AM NOT DONE
//
// mul перемножает две матрицы 2×2.
// Тренирует: тройной цикл по массивам.
// Сложность: medium
package main_test

import "testing"

func mul(a, b [2][2]int) [2][2]int {
	var c [2][2]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = a[i][j] * b[i][j]
		}
	}
	return c
}

func TestMul(t *testing.T) {
	a := [2][2]int{{1, 2}, {3, 4}}
	b := [2][2]int{{5, 6}, {7, 8}}
	if got := mul(a, b); got != [2][2]int{{19, 22}, {43, 50}} {
		t.Errorf("mul = %v", got)
	}
}
