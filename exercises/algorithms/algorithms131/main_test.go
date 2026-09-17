// algorithms131
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: сложение в столбик. Сложите два двоичных числа, заданных строками,
// и верните результат строкой без ведущих нулей (кроме самого "0").
// Сложность: easy. Ожидаемая асимптотика: O(n) по времени, O(n) по памяти
package main_test

import "testing"

func addBinary(a, b string) string {
	return ""
}

func TestAddBinary(t *testing.T) {
	cases := []struct{ a, b, want string }{
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
		{"0", "0", "0"},
		{"", "101", "101"},
		{"1111", "1111", "11110"},
	}
	for _, c := range cases {
		if got := addBinary(c.a, c.b); got != c.want {
			t.Errorf("addBinary(%q, %q) = %q, want %q", c.a, c.b, got, c.want)
		}
	}
}
