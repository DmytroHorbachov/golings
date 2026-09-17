// if51
// Make the tests pass!

// I AM NOT DONE
//
// sameRoute сравнивает два маршрута (срезы названий остановок).
// Код не компилируется.
// Тренирует: срезы сравнимы только с nil.
// Сложность: hard
package main_test

import "testing"

func sameRoute(a, b []string) bool {
	if a == b {
		return true
	}
	return false
}

func TestSameRoute(t *testing.T) {
	if !sameRoute([]string{"A", "B"}, []string{"A", "B"}) {
		t.Errorf("identical routes should match")
	}
	if sameRoute([]string{"A", "B"}, []string{"B", "A"}) || sameRoute([]string{"A"}, nil) {
		t.Errorf("different routes should not match")
	}
	if !sameRoute(nil, []string{}) {
		t.Errorf("two empty routes should match")
	}
}
