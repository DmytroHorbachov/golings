// range18
// Make the tests pass!

// I AM NOT DONE
//
// oldest возвращает возраст самого старшего человека.
// Тренирует: range по срезу структур и поиск максимума.
// Сложность: easy
package main_test

import "testing"

type Person struct {
	Name string
	Age  int
}

func oldest(ps []Person) int {
	max := 0
	for _, p := range ps {
		if p.Age < max {
			max = p.Age
		}
	}
	return max
}

func TestOldest(t *testing.T) {
	if got := oldest([]Person{{"a", 30}, {"b", 71}, {"c", 45}}); got != 71 {
		t.Errorf("oldest = %d, want 71", got)
	}
}
