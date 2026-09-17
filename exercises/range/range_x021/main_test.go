// range_x021: Чередование по индексу
// Make the tests pass!
// I AM NOT DONE
//
// stripes возвращает "#" для чётных позиций и "." для нечётных.
// Тренирует: использование индекса range в вычислениях.
// Сложность: easy
package main_test

import (
	"strings"
	"testing"
)

func stripes(n int) string {
	var b strings.Builder
	for i := range make([]struct{}, n) {
		if i%2 == 1 {
			b.WriteString("#")
		} else {
			b.WriteString(".")
		}
	}
	return b.String()
}

func TestStripes(t *testing.T) {
	if got := stripes(5); got != "#.#.#" {
		t.Errorf("stripes(5) = %q", got)
	}
}
