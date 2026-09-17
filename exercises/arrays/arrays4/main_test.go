// arrays4
// Make the tests pass!

// I AM NOT DONE
//
// Названия цветов хранятся в массиве размера numColors. Добавили цвет Purple,
// но размер массива задан числом и тест падает.
// Тренирует: связывание размера массива с константой-счётчиком iota.
// Сложность: hard
package main_test

import "testing"

type Color int

const (
	Red Color = iota
	Green
	Blue
	Purple
)

var colorNames = [3]string{"red", "green", "blue"}

func (c Color) String() string {
	if c < 0 || int(c) >= len(colorNames) {
		return "unknown"
	}
	return colorNames[c]
}

func TestColorNames(t *testing.T) {
	if Purple.String() != "purple" || Red.String() != "red" || Color(42).String() != "unknown" {
		t.Errorf("names: %s %s %s", Purple, Red, Color(42))
	}
}
