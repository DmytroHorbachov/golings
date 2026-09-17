// range71
// Make the tests pass!

// I AM NOT DONE
//
// visit обходит элементы через колбэк и должна прекратить обход на "end".
// Код не компилируется: break внутри функционального литерала не относится к циклу.
// Тренирует: break допустим только внутри for/switch/select той же функции.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

func walk(items []string, f func(string) bool) {
	for _, it := range items {
		if !f(it) {
			return
		}
	}
}

func visit(items []string) []string {
	var seen []string
	walk(items, func(s string) bool {
		if s == "end" {
			break
		}
		seen = append(seen, s)
		return true
	})
	return seen
}

func TestVisit(t *testing.T) {
	if got := visit([]string{"a", "b", "end", "c"}); !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Errorf("visit = %v", got)
	}
}
