// if_x085: Последний символ
// Make the tests pass!
// I AM NOT DONE
//
// isQuestion должна вернуть true для строк, оканчивающихся на '?'.
// На пустой строке функция паникует.
// Тренирует: len(s)-1 для пустой строки равно -1.
// Сложность: hard
package main_test

import "testing"

func isQuestion(s string) bool {
	if s[len(s)-1] == '?' {
		return true
	}
	return false
}

func TestIsQuestion(t *testing.T) {
	cases := map[string]bool{"why?": true, "ok": false, "": false, "?": true}
	for in, want := range cases {
		if got := isQuestion(in); got != want {
			t.Errorf("isQuestion(%q) = %v, want %v", in, got, want)
		}
	}
}
