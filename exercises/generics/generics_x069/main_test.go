// generics_x069: < без упорядочивания
// Make the tests pass!
// I AM NOT DONE
//
// MaxOf возвращает максимум среза. Код не компилируется:
// comparable не поддерживает <.
// Тренирует: comparable даёт только == и !=.
// Сложность: hard
package main_test

import "testing"

func MaxOf[T comparable](s []T) T {
	m := s[0]
	for _, v := range s[1:] {
		if v > m {
			m = v
		}
	}
	return m
}

func TestMaxOf(t *testing.T) {
	if MaxOf([]int{3, 9, 2}) != 9 || MaxOf([]string{"b", "c", "a"}) != "c" {
		t.Errorf("MaxOf works incorrectly")
	}
}
