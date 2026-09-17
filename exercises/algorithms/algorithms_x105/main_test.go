// algorithms_x105: Word Break (разбиение строки на слова)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: одномерное ДП по позициям строки. Можно ли разбить строку
// на последовательность слов из словаря (слова можно использовать много раз)?
// Сложность: medium. Ожидаемая асимптотика: O(n²·L) по времени, O(n) по памяти
package main_test

import "testing"

func wordBreak(s string, dict []string) bool {
	return false
}

func TestWordBreak(t *testing.T) {
	cases := []struct {
		s    string
		dict []string
		want bool
	}{
		{"leetcode", []string{"leet", "code"}, true},
		{"applepenapple", []string{"apple", "pen"}, true},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
		{"", []string{"a"}, true},
		{"a", nil, false},
		{"aaaaaaa", []string{"aaa", "aaaa"}, true},
	}
	for _, c := range cases {
		if got := wordBreak(c.s, c.dict); got != c.want {
			t.Errorf("wordBreak(%q, %v) = %v, want %v", c.s, c.dict, got, c.want)
		}
	}
}
