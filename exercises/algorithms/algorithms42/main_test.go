// algorithms42
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: система счисления по основанию 26. Переведите заголовок столбца
// ("A", "B", ..., "Z", "AA", ...) в его номер.
// Сложность: easy. Ожидаемая асимптотика: O(n) по времени, O(1) по памяти
package main_test

import "testing"

func titleToNumber(s string) int {
	return 0
}

func TestTitleToNumber(t *testing.T) {
	cases := map[string]int{"A": 1, "AB": 28, "ZY": 701, "": 0, "FXSHRXW": 2147483647}
	for in, want := range cases {
		if got := titleToNumber(in); got != want {
			t.Errorf("titleToNumber(%q) = %d, want %d", in, got, want)
		}
	}
}
