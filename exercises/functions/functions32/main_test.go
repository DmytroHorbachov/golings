// functions32
// Make the tests pass!

// I AM NOT DONE
//
// reverse должна рекурсивно развернуть строку из рун.
// Тренирует: рекурсию на срезе рун и сборку результата.
// Сложность: medium
package main_test

import "testing"

func reverse(s string) string {
	r := []rune(s)
	if len(r) <= 1 {
		return s
	}
	return string(r[0]) + reverse(string(r[1:]))
}

func reverseAll(items []string) []string {
	out := make([]string, len(items))
	for i, s := range items {
		out[i] = s
	}
	return out
}

func TestReverseAll(t *testing.T) {
	got := reverseAll([]string{"abc", "мир", ""})
	if got[0] != "cba" || got[1] != "рим" || got[2] != "" {
		t.Errorf("reverseAll = %q", got)
	}
}
