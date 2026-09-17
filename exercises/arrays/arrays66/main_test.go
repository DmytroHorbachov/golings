// arrays66
// Make the tests pass!

// I AM NOT DONE
//
// Тип Palette — массив цветов. Метод Set должен менять цвет по индексу.
// После вызова палитра остаётся прежней.
// Тренирует: методы с получателем-значением работают с копией массива.
// Сложность: hard
package main_test

import "testing"

type Palette [3]string

func (p Palette) Set(i int, c string) {
	p[i] = c
}

func TestPaletteSet(t *testing.T) {
	p := Palette{"red", "green", "blue"}
	p.Set(1, "yellow")
	if p[1] != "yellow" {
		t.Errorf("p[1] = %s, want yellow", p[1])
	}
}
