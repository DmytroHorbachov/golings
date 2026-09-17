// anonymous_functions59
// Make the tests pass!

// I AM NOT DONE
//
// Рекурсивный литерал объявлен через var, но вызывается раньше, чем ему
// присвоено значение, и программа паникует.
// Тренирует: нулевое значение функциональной переменной — nil.
// Сложность: hard
package main_test

import "testing"

func depth(tree map[string][]string, root string) int {
	var measure func(string) int
	result := measure(root)
	measure = func(n string) int {
		best := 0
		for _, c := range tree[n] {
			if d := measure(c); d > best {
				best = d
			}
		}
		return best + 1
	}
	return result
}

func TestDepth(t *testing.T) {
	tree := map[string][]string{"a": {"b", "c"}, "c": {"d"}}
	if got := depth(tree, "a"); got != 3 {
		t.Errorf("depth = %d, want 3", got)
	}
}
