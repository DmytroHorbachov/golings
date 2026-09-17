// algorithms_x140: Happy Number (счастливое число)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: обнаружение цикла (Флойд). Число счастливое, если повторная замена
// на сумму квадратов его цифр приводит к 1.
// Сложность: easy. Ожидаемая асимптотика: O(log n) по времени, O(1) по памяти
package main_test

import "testing"

func isHappy(n int) bool {
	return false
}

func TestIsHappy(t *testing.T) {
	cases := map[int]bool{19: true, 2: false, 1: true, 7: true, 116: false}
	for in, want := range cases {
		if got := isHappy(in); got != want {
			t.Errorf("isHappy(%d) = %v, want %v", in, got, want)
		}
	}
}
