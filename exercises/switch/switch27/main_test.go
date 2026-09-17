// switch27
// Make the tests pass!

// I AM NOT DONE
//
// pieceValue возвращает ценность шахматной фигуры: пешка 1, конь и слон 3,
// ладья 5, ферзь 9.
// Тренирует: объединение значений в одной ветке.
// Сложность: easy
package main_test

import "testing"

func pieceValue(p byte) int {
	switch p {
	case 'P':
		return 1
	case 'N':
		return 3
	case 'R':
		return 5
	case 'Q':
		return 9
	}
	return 0
}

func TestPieceValue(t *testing.T) {
	cases := map[byte]int{'P': 1, 'N': 3, 'B': 3, 'R': 5, 'Q': 9, 'K': 0}
	for in, want := range cases {
		if got := pieceValue(in); got != want {
			t.Errorf("pieceValue(%c) = %d, want %d", in, got, want)
		}
	}
}
