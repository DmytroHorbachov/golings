// structs42
// Make the tests pass!

// I AM NOT DONE
//
// newColor создаёт цвет позиционным литералом (порядок полей R, G, B).
// Тренирует: литерал структуры без имён полей.
// Сложность: easy
package main_test

import "testing"

type Color struct{ R, G, B uint8 }

func orange() Color {
	return Color{0, 165, 255}
}

func TestOrange(t *testing.T) {
	if c := orange(); c.R != 255 || c.G != 165 || c.B != 0 {
		t.Errorf("orange = %+v", c)
	}
}
