// range54
// Make the tests pass!

// I AM NOT DONE
//
// countByte считает вхождения заданного байта в строку.
// Код не компилируется: range по строке выдаёт rune, а искомое значение — byte.
// Тренирует: типы значений range по строке и по []byte различаются.
// Сложность: hard
package main_test

import "testing"

func countByte(s string, c byte) int {
	n := 0
	for _, r := range s {
		if r == c {
			n++
		}
	}
	return n
}

func TestCountByte(t *testing.T) {
	if got := countByte("a,b,,c", ','); got != 3 {
		t.Errorf("countByte = %d, want 3", got)
	}
}
