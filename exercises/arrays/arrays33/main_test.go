// arrays33
// Make the tests pass!

// I AM NOT DONE
//
// setCell должна поставить отметку в клетку поля.
// Строка извлекается в переменную, изменяется, но поле остаётся прежним.
// Тренирует: присваивание элемента-массива копирует его.
// Сложность: hard
package main_test

import "testing"

type Grid [3][3]byte

func (g *Grid) setCell(r, c int, v byte) {
	row := g[r]
	row[c] = v
}

func TestSetCell(t *testing.T) {
	var g Grid
	g.setCell(1, 2, 'X')
	if g[1][2] != 'X' {
		t.Errorf("g[1][2] = %q, want 'X'", g[1][2])
	}
}
