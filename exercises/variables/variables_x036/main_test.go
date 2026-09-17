// variables_x036: iota в одной строке
// Make the tests pass!
// I AM NOT DONE
//
// Ожидается, что Width=0, Height=1, Depth=2.
// Сейчас две константы объявлены в одной строке спецификации.
// Тренирует: iota увеличивается по строкам (ConstSpec), а не по именам.
// Сложность: hard
package main_test

import "testing"

const (
	Width, Height = iota, iota
	Depth         = iota
)

func TestAxes(t *testing.T) {
	if Width != 0 || Height != 1 || Depth != 2 {
		t.Errorf("Width=%d Height=%d Depth=%d, want 0 1 2", Width, Height, Depth)
	}
}
