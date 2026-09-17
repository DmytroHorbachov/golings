// variables_x077: Шестнадцатеричный литерал
// Make the tests pass!
// I AM NOT DONE
//
// Константа White должна быть равна 0xFFFFFF (16777215).
// Тренирует: шестнадцатеричные целочисленные литералы.
// Сложность: easy
package main_test

import "testing"

const White = 0xFFFF

func TestWhite(t *testing.T) {
	if White != 16777215 {
		t.Errorf("White = %d, want 16777215", White)
	}
}
