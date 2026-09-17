// slices104
// Make the tests pass!

// I AM NOT DONE
//
// addRow добавляет строку в таблицу [][]string.
// Тренирует: append среза в срез срезов.
// Сложность: easy
package main_test

import "testing"

func addRow(table [][]string, row []string) [][]string {
	return append(table, row[:1])
}

func TestAddRow(t *testing.T) {
	tb := addRow([][]string{{"id", "name"}}, []string{"1", "ann"})
	if len(tb) != 2 || len(tb[1]) != 2 || tb[1][1] != "ann" {
		t.Errorf("addRow = %v", tb)
	}
}
