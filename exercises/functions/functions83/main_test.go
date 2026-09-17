// functions83
// Make the tests pass!

// I AM NOT DONE
//
// place должна поставить отметку на игровом поле.
// После вызовов поле остаётся пустым.
// Тренирует: массивы (в отличие от срезов) копируются при передаче в функцию.
// Сложность: hard
package main_test

import "testing"

type Board [3][3]rune

func place(b Board, r, c int, mark rune) {
	b[r][c] = mark
}

func play() Board {
	var b Board
	place(b, 0, 0, 'X')
	place(b, 1, 1, 'O')
	return b
}

func TestPlay(t *testing.T) {
	b := play()
	if b[0][0] != 'X' || b[1][1] != 'O' {
		t.Errorf("board = %q, want X at (0,0) and O at (1,1)", b)
	}
}
