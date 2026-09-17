// arrays_x079: Метод возвращает копию
// Make the tests pass!
// I AM NOT DONE
//
// Board.Row возвращает строку поля, а вызывающий код хочет её изменить.
// Изменения теряются, потому что возвращается копия массива.
// Тренирует: возврат массива из функции — это копирование.
// Сложность: hard
package main_test

import "testing"

type Board struct {
	cells [2][4]int
}

func (b *Board) Row(i int) [4]int {
	return b.cells[i]
}

func TestRow(t *testing.T) {
	var b Board
	row := b.Row(1)
	row[3] = 42
	if b.cells[1][3] != 42 {
		t.Errorf("cells[1][3] = %d, want 42", b.cells[1][3])
	}
}
