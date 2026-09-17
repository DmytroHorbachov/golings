// switch8
// Make the tests pass!

// I AM NOT DONE
//
// heading определяет уровень заголовка Markdown ("# " — 1, "## " — 2, "### " — 3),
// для остальных строк — 0.
// Тренирует: подсчёт перед switch и проверку деталей в ветке.
// Сложность: medium
package main_test

import "testing"

func heading(line string) int {
	n := 0
	for n < len(line) && line[n] == '#' {
		n++
	}
	switch n {
	case 1, 2, 3:
		return n
	}
	return 0
}

func TestHeading(t *testing.T) {
	cases := map[string]int{"# Title": 1, "### Sub": 3, "#### Deep": 0, "#tag": 0, "text": 0, "##": 0}
	for in, want := range cases {
		if got := heading(in); got != want {
			t.Errorf("heading(%q) = %d, want %d", in, got, want)
		}
	}
}
