// range_x096: Результат TrimSpace
// Make the tests pass!
// I AM NOT DONE
//
// trimAll убирает пробелы по краям каждой строки среза.
// Строки не меняются: результат функции отбрасывается.
// Тренирует: строки неизменяемы, функции strings возвращают новые строки.
// Сложность: hard
package main_test

import (
	"reflect"
	"strings"
	"testing"
)

func trimAll(lines []string) {
	for _, l := range lines {
		l = strings.TrimSpace(l)
		_ = l
	}
}

func TestTrimAll(t *testing.T) {
	lines := []string{"  a ", "b\n", "c"}
	trimAll(lines)
	if !reflect.DeepEqual(lines, []string{"a", "b", "c"}) {
		t.Errorf("trimAll = %q", lines)
	}
}
