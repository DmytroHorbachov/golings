// variables_x014: Пропуск значений iota
// Make the tests pass!
// I AM NOT DONE
//
// Коды ответа должны быть: OK=0, (1 пропущен), NotFound=2, Forbidden=3.
// Тренирует: пропуск значений iota с помощью _.
// Сложность: easy
package main_test

import "testing"

const (
	OK = iota
	NotFound
	Forbidden
)

func TestCodes(t *testing.T) {
	if OK != 0 || NotFound != 2 || Forbidden != 3 {
		t.Errorf("OK=%d NotFound=%d Forbidden=%d, want 0 2 3", OK, NotFound, Forbidden)
	}
}
