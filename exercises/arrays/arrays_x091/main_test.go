// arrays_x091: Размер из переменной
// Make the tests pass!
// I AM NOT DONE
//
// table должна создать таблицу размера n×n. Код не компилируется:
// размер массива должен быть константой.
// Тренирует: для размеров, известных только во время выполнения, нужны срезы.
// Сложность: hard
package main_test

import "testing"

func table(n int) [][]int {
	var t [n][n]int
	out := make([][]int, n)
	for i := range t {
		out[i] = t[i][:]
	}
	return out
}

func TestTable(t *testing.T) {
	tb := table(3)
	if len(tb) != 3 || len(tb[2]) != 3 || tb[2][2] != 9 || tb[0][1] != 2 {
		t.Errorf("table(3) = %v", tb)
	}
}
