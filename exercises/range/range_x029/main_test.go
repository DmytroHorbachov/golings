// range_x029: Множество
// Make the tests pass!
// I AM NOT DONE
//
// countTrue считает включённые флаги.
// Тренирует: range по map[string]bool.
// Сложность: easy
package main_test

import "testing"

func countTrue(flags map[string]bool) int {
	n := 0
	for _, on := range flags {
		if !on {
			n++
		}
	}
	return n
}

func TestCountTrue(t *testing.T) {
	if got := countTrue(map[string]bool{"a": true, "b": false, "c": true}); got != 2 {
		t.Errorf("countTrue = %d, want 2", got)
	}
}
