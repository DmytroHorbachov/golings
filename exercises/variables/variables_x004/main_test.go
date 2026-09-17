// variables_x004: iota с единицы
// Make the tests pass!
// I AM NOT DONE
//
// Уровни логирования должны нумероваться с единицы: Debug=1, Info=2, Warn=3.
// Тренирует: константы и генератор iota.
// Сложность: easy
package main_test

import "testing"

const (
	Debug = iota
	Info
	Warn
)

func TestLevels(t *testing.T) {
	if Debug != 1 || Info != 2 || Warn != 3 {
		t.Errorf("got Debug=%d Info=%d Warn=%d, want 1 2 3", Debug, Info, Warn)
	}
}
