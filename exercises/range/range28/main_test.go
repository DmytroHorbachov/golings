// range28
// Make the tests pass!

// I AM NOT DONE
//
// takeUntilStop должна обработать элементы до "stop" и вернуть обработанные.
// Срез обрезается внутри цикла, но цикл продолжается.
// Тренирует: range вычисляет срез один раз; переприсваивание переменной его не меняет.
// Сложность: hard
package main_test

import (
	"reflect"
	"strings"
	"testing"
)

func takeUntilStop(s []string) []string {
	var out []string
	for i, v := range s {
		if v == "stop" {
			s = s[:i]
			continue
		}
		_ = i
		out = append(out, strings.ToUpper(v))
	}
	return out
}

func TestTakeUntilStop(t *testing.T) {
	got := takeUntilStop([]string{"a", "b", "stop", "c"})
	if !reflect.DeepEqual(got, []string{"A", "B"}) {
		t.Errorf("takeUntilStop = %v", got)
	}
}
