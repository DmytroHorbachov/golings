// algorithms94
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: стек индексов. Найдите длину самой длинной подстроки из '(' и ')',
// являющейся правильной скобочной последовательностью.
// Сложность: hard. Ожидаемая асимптотика: O(n) по времени, O(n) по памяти
package main_test

import "testing"

func longestValidParentheses(s string) int {
	return 0
}

func TestLongestValidParentheses(t *testing.T) {
	cases := map[string]int{"(()": 2, ")()())": 4, "": 0, "(": 0, "()(())": 6, "())(()": 2, "(()())": 6}
	for in, want := range cases {
		if got := longestValidParentheses(in); got != want {
			t.Errorf("longestValidParentheses(%q) = %d, want %d", in, got, want)
		}
	}
}
