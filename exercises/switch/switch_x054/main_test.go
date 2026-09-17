// switch_x054: Сравнение словами
// Make the tests pass!
// I AM NOT DONE
//
// compareWords сравнивает два числа и возвращает "less", "equal" или "greater".
// Тренирует: switch без тега с тремя исходами.
// Сложность: medium
package main_test

import "testing"

func compareWords(a, b int) string {
	switch {
	case a <= b:
		return "less"
	case a >= b:
		return "greater"
	}
	return "equal"
}

func TestCompareWords(t *testing.T) {
	cases := [][2]int{{1, 2}, {2, 2}, {3, 2}}
	want := []string{"less", "equal", "greater"}
	for i, c := range cases {
		if got := compareWords(c[0], c[1]); got != want[i] {
			t.Errorf("compareWords(%d, %d) = %s, want %s", c[0], c[1], got, want[i])
		}
	}
}
