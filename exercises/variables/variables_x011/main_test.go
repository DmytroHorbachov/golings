// variables_x011: Разделители разрядов
// Make the tests pass!
// I AM NOT DONE
//
// Константа Million должна быть равна одному миллиону.
// Кто-то ошибся при наборе числа, и цифр не хватает.
// Тренирует: числовые литералы с разделителем _.
// Сложность: easy
package main_test

import "testing"

const Million = 1_000_00

func TestMillion(t *testing.T) {
	if Million != 1000000 {
		t.Errorf("Million = %d, want 1000000", Million)
	}
}
