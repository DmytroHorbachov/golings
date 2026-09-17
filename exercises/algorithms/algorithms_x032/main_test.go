// algorithms_x032: Basic Calculator (калькулятор со скобками)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: стек. Вычислите выражение из целых чисел, '+', '-', скобок и пробелов.
// Унарный минус допускается перед числом или скобкой.
// Сложность: hard. Ожидаемая асимптотика: O(n) по времени, O(n) по памяти
package main_test

import "testing"

func calculate(s string) int {
	return 0
}

func TestCalculate(t *testing.T) {
	cases := map[string]int{
		"1 + 1":               2,
		" 2-1 + 2 ":           3,
		"(1+(4+5+2)-3)+(6+8)": 23,
		"-(2+3)":              -5,
		"42":                  42,
		"- (3 - (4 - 10))":    -9,
		"2147483647 + 1":      2147483648,
	}
	for in, want := range cases {
		if got := calculate(in); got != want {
			t.Errorf("calculate(%q) = %d, want %d", in, got, want)
		}
	}
}
