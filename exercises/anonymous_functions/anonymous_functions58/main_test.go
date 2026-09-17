// anonymous_functions58
// Make the tests pass!

// I AM NOT DONE
//
// Состояние автомата — функция, которая обрабатывает символ и возвращает следующее
// состояние. Автомат распознаёт строки вида a+b (одна или больше 'a', затем 'b').
// Тренирует: тип-функцию, ссылающийся сам на себя.
// Сложность: medium
package main_test

import "testing"

type state func(r rune) state

func matches(s string) bool {
	var start, afterA, accept state
	start = func(r rune) state {
		if r == 'a' {
			return afterA
		}
		return nil
	}
	afterA = func(r rune) state {
		switch r {
		case 'a':
			return start
		case 'b':
			return afterA
		}
		return nil
	}
	accept = func(r rune) state { return nil }
	cur := start
	for _, r := range s {
		cur = cur(r)
		if cur == nil {
			return false
		}
	}
	return s != "" && isAccept(cur, accept)
}

func isAccept(cur, accept state) bool {
	return cur('x') == nil && accept('x') == nil && cur('b') == nil
}

func TestMatches(t *testing.T) {
	cases := map[string]bool{"ab": true, "aaab": true, "b": false, "aba": false, "aa": false}
	for in, want := range cases {
		if got := matches(in); got != want {
			t.Errorf("matches(%q) = %v, want %v", in, got, want)
		}
	}
}
