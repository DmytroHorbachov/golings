// arrays_x090: len в константе
// Make the tests pass!
// I AM NOT DONE
//
// Размер буфера должен вычисляться на этапе компиляции из количества заголовков.
// Код не компилируется: len от среза — не константа.
// Тренирует: len от массива — константное выражение, от среза — нет.
// Сложность: hard
package main_test

import "testing"

var headers = []string{"id", "name", "email"}

const numHeaders = len(headers)

var widths [numHeaders]int

func TestNumHeaders(t *testing.T) {
	if numHeaders != 3 || len(widths) != 3 {
		t.Errorf("numHeaders = %d", numHeaders)
	}
}
