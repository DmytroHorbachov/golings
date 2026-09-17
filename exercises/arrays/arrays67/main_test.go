// arrays67
// Make the tests pass!

// I AM NOT DONE
//
// board строит доску 8×8, где true — чёрная клетка; клетка a1 ([7][0]) чёрная.
// Тренирует: заполнение двумерного массива по формуле.
// Сложность: medium
package main_test

import "testing"

func board() [8][8]bool {
	var b [8][8]bool
	for r := range b {
		b[r][0] = r%2 == 0
	}
	return b
}

func TestBoard(t *testing.T) {
	b := board()
	if !b[7][0] || b[7][1] || b[0][0] || !b[0][1] || !b[3][4] {
		t.Errorf("board colors are wrong")
	}
}
