// algorithms_x095: Alien Dictionary (порядок букв инопланетян)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: топологическая сортировка по парам соседних слов. По списку слов,
// отсортированному в неизвестном алфавите, восстановите порядок букв.
// При неоднозначности берите меньшую по коду букву; при противоречии — "".
// Сложность: hard. Ожидаемая асимптотика: O(C) по времени, O(1) по памяти
package main_test

import (
	"sort"
	"testing"
)

func alienOrder(words []string) string {
	return ""
}

func TestAlienOrder(t *testing.T) {
	_ = sort.Ints
	cases := []struct {
		words []string
		want  string
	}{
		{[]string{"wrt", "wrf", "er", "ett", "rftt"}, "wertf"},
		{[]string{"z", "x"}, "zx"},
		{[]string{"z", "x", "z"}, ""},
		{[]string{"abc", "ab"}, ""},
		{[]string{"ab", "adc"}, "abcd"},
	}
	for _, c := range cases {
		if got := alienOrder(c.words); got != c.want {
			t.Errorf("alienOrder(%v) = %q, want %q", c.words, got, c.want)
		}
	}
}
