// slices38
// Make the tests pass!

// I AM NOT DONE
//
// makeBoard строит поле rows×cols из точек. После изменения одной клетки
// меняется весь столбец.
// Тренирует: одна и та же строка-срез, добавленная несколько раз, — это одна строка.
// Сложность: hard
package main_test

import "testing"

func makeBoard(rows, cols int) [][]byte {
	board := make([][]byte, 0, rows)
	row := make([]byte, cols)
	for c := range row {
		row[c] = '.'
	}
	for r := 0; r < rows; r++ {
		board = append(board, row)
	}
	return board
}

func TestMakeBoard(t *testing.T) {
	b := makeBoard(3, 3)
	b[0][1] = 'X'
	if b[1][1] != '.' || b[2][1] != '.' || b[0][1] != 'X' {
		t.Errorf("board = %q", b)
	}
}
