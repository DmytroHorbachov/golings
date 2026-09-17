// primitive_types_x071: Индекс в range по строке
// Make the tests pass!
// I AM NOT DONE
//
// charAt должна вернуть n-й символ строки (считая с нуля) или пустую строку.
// range по строке выдаёт байтовые смещения, а не номера символов.
// Тренирует: индекс в for i, r := range s — это позиция байта.
// Сложность: hard
package main_test

import "testing"

func charAt(s string, n int) string {
	for i, r := range s {
		if i == n {
			return string(r)
		}
	}
	return ""
}

func TestCharAt(t *testing.T) {
	cases := []struct {
		s    string
		n    int
		want string
	}{{"hello", 1, "e"}, {"мир", 1, "и"}, {"мир", 2, "р"}, {"日本語", 2, "語"}, {"go", 5, ""}}
	for _, c := range cases {
		if got := charAt(c.s, c.n); got != c.want {
			t.Errorf("charAt(%q, %d) = %q, want %q", c.s, c.n, got, c.want)
		}
	}
}
