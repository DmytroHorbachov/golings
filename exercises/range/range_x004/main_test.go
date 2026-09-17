// range_x004: Значения map
// Make the tests pass!
// I AM NOT DONE
//
// totalAge складывает возрасты всех людей из map.
// Тренирует: range по map с двумя переменными.
// Сложность: easy
package main_test

import "testing"

func totalAge(ages map[string]int) int {
	total := 0
	for name, age := range ages {
		_ = name
		total += len(name)
	}
	return total
}

func TestTotalAge(t *testing.T) {
	if got := totalAge(map[string]int{"ann": 30, "bob": 25}); got != 55 {
		t.Errorf("totalAge = %d, want 55", got)
	}
}
