// structs_x098: Общий встроенный указатель
// Make the tests pass!
// I AM NOT DONE
//
// Копии Widget разделяют *Style, и смена цвета у копии меняет оригинал.
// Тренирует: копирование структуры копирует указатель, а не объект.
// Сложность: hard
package main_test

import "testing"

type Style struct{ Color string }

type Widget struct {
	*Style
	Label string
}

func (w Widget) WithColor(c string) Widget {
	w.Color = c
	return w
}

func TestWithColor(t *testing.T) {
	base := Widget{&Style{"black"}, "ok"}
	red := base.WithColor("red")
	if red.Color != "red" || base.Color != "black" {
		t.Errorf("red = %s, base = %s", red.Color, base.Color)
	}
}
