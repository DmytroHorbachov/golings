// functions_x021: Результаты как аргументы
// Make the tests pass!
// I AM NOT DONE
//
// Функция divmod возвращает два значения, которые сразу передаются в format.
// Код не компилируется из-за лишнего аргумента.
// Тренирует: передачу нескольких результатов напрямую в другую функцию.
// Сложность: easy
package main_test

import (
	"strconv"
	"testing"
)

func divmod(a, b int) (int, int) { return a / b, a % b }

func format(q, r int) string {
	return strconv.Itoa(q) + " r" + strconv.Itoa(r)
}

func describe(a, b int) string {
	return format(divmod(a, b), 0)
}

func TestDescribe(t *testing.T) {
	if got := describe(17, 5); got != "3 r2" {
		t.Errorf("describe(17, 5) = %q, want %q", got, "3 r2")
	}
}
