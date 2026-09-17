// range94
// Make the tests pass!

// I AM NOT DONE
//
// withCommas ставит запятую после каждого символа, кроме последнего.
// Если последний символ многобайтовый, запятая появляется и после него.
// Тренирует: i == len(s)-1 не определяет последнюю руну.
// Сложность: hard
package main_test

import (
	"strings"
	"testing"
	"unicode/utf8"
)

func withCommas(s string) string {
	var b strings.Builder
	for i, r := range s {
		b.WriteRune(r)
		if i != len(s)-1 {
			b.WriteByte(',')
		}
	}
	return b.String()
}

func TestWithCommas(t *testing.T) {
	_ = utf8.RuneLen
	cases := map[string]string{"abc": "a,b,c", "abя": "a,b,я", "ё": "ё"}
	for in, want := range cases {
		if got := withCommas(in); got != want {
			t.Errorf("withCommas(%q) = %q, want %q", in, got, want)
		}
	}
}
