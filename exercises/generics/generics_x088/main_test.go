// generics_x088: Потеря именованного типа среза
// Make the tests pass!
// I AM NOT DONE
//
// Filter возвращает []T, и результат теряет методы именованного типа Names.
// Код не компилируется: у []string нет метода Join.
// Тренирует: паттерн S ~[]E сохраняет тип среза.
// Сложность: hard
package main_test

import (
	"strings"
	"testing"
)

type Names []string

func (n Names) Join() string { return strings.Join(n, ",") }

func Filter[E any](s []E, keep func(E) bool) []E {
	var out []E
	for _, v := range s {
		if keep(v) {
			out = append(out, v)
		}
	}
	return out
}

func TestFilterNames(t *testing.T) {
	names := Names{"ann", "", "bob"}
	got := Filter(names, func(s string) bool { return s != "" }).Join()
	if got != "ann,bob" {
		t.Errorf("got %q", got)
	}
}
