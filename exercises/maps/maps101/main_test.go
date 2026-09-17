// maps101
// Make the tests pass!

// I AM NOT DONE
//
// clearAll удаляет все элементы, сохраняя саму map (другие ссылаются на неё).
// Тренирует: delete во время range безопасен.
// Сложность: easy
package main_test

import "testing"

func clearAll(m map[string]int) {
	for k := range m {
		m[k] = 0
	}
}

func TestClearAll(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	alias := m
	clearAll(m)
	if len(alias) != 0 {
		t.Errorf("alias = %v, want empty", alias)
	}
}
