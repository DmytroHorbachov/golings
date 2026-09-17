// slices_x089: nil-строки сетки
// Make the tests pass!
// I AM NOT DONE
//
// setCell пишет значение в разреженную сетку, где строки создаются по требованию.
// Запись в ещё не созданную строку паникует.
// Тренирует: make([][]T, n) создаёт n nil-срезов.
// Сложность: hard
package main_test

import "testing"

type Sparse struct {
	rows [][]int
	cols int
}

func newSparse(r, c int) *Sparse {
	return &Sparse{rows: make([][]int, r), cols: c}
}

func (s *Sparse) setCell(r, c, v int) {
	s.rows[r][c] = v
}

func TestSetCell(t *testing.T) {
	s := newSparse(100, 100)
	s.setCell(42, 7, 1)
	s.setCell(42, 8, 2)
	if s.rows[42][7] != 1 || s.rows[42][8] != 2 || s.rows[0] != nil {
		t.Errorf("sparse grid is wrong")
	}
}
