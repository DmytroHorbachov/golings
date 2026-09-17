// generics62
// Make the tests pass!

// I AM NOT DONE
//
// Last[S []E, E any] не принимает именованный тип среза.
// Код не компилируется при вызове с Stack.
// Тренирует: []E в ограничении без ~ соответствует только самому []E.
// Сложность: hard
package main_test

import "testing"

type Stack []int

func Last[S []E, E any](s S) E {
	return s[len(s)-1]
}

func TestLast(t *testing.T) {
	if Last(Stack{1, 2, 3}) != 3 || Last([]string{"a"}) != "a" {
		t.Errorf("Last works incorrectly")
	}
}
