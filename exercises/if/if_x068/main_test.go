// if_x068: Порядок в &&
// Make the tests pass!
// I AM NOT DONE
//
// isComment должна вернуть true для строк, начинающихся с '#'.
// На пустой строке функция паникует.
// Тренирует: сокращённое вычисление (short-circuit) логических операторов.
// Сложность: hard
package main_test

import "testing"

func isComment(line string) bool {
	if line[0] == '#' && len(line) > 0 {
		return true
	}
	return false
}

func TestIsComment(t *testing.T) {
	cases := map[string]bool{"# note": true, "code": false, "": false, "#": true}
	for in, want := range cases {
		if got := isComment(in); got != want {
			t.Errorf("isComment(%q) = %v, want %v", in, got, want)
		}
	}
}
