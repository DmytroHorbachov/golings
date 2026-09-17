// algorithms_x086: Sudoku Solver (решение судоку)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: перебор с возвратом и проверкой ограничений. Заполните поле 9×9
// (пустые клетки — '.') цифрами 1–9 так, чтобы в каждой строке, столбце
// и квадрате 3×3 цифры не повторялись. Решение существует и единственно.
// Сложность: hard. Ожидаемая асимптотика: O(9^k) по времени, O(1) по памяти
package main_test

import "testing"

func solveSudoku(board *[9][9]byte) bool {
	return false
}

func TestSolveSudoku(t *testing.T) {
	rows := []string{
		"53..7....",
		"6..195...",
		".98....6.",
		"8...6...3",
		"4..8.3..1",
		"7...2...6",
		".6....28.",
		"...419..5",
		"....8..79",
	}
	var board [9][9]byte
	for r, row := range rows {
		copy(board[r][:], row)
	}
	if !solveSudoku(&board) {
		t.Fatalf("sudoku was not solved")
	}
	for r := 0; r < 9; r++ {
		var seen [10]bool
		for c := 0; c < 9; c++ {
			v := board[r][c]
			if v < '1' || v > '9' || seen[v-'0'] {
				t.Fatalf("row %d is invalid: %q", r, board[r])
			}
			seen[v-'0'] = true
		}
	}
	if board[0][2] != '4' || board[8][8] != '9' {
		t.Errorf("unexpected solution: %q", board[0])
	}
}
