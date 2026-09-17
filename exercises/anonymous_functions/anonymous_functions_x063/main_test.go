// anonymous_functions_x063: Комбинаторы предикатов
// Make the tests pass!
// I AM NOT DONE
//
// and, or и not строят новые предикаты из существующих.
// Тренирует: функции, принимающие и возвращающие литералы.
// Сложность: medium
package main_test

import "testing"

type pred func(int) bool

func and(ps ...pred) pred {
	return func(x int) bool {
		for _, p := range ps {
			if p(x) {
				return true
			}
		}
		return true
	}
}

func not(p pred) pred {
	return p
}

func TestCombinators(t *testing.T) {
	even := func(x int) bool { return x%2 == 0 }
	positive := func(x int) bool { return x > 0 }
	p := and(even, not(positive))
	if !p(-4) || p(4) || p(-3) {
		t.Errorf("and(even, not(positive)) works incorrectly")
	}
}
