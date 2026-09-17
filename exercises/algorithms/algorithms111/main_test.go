// algorithms111
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: битовые операции. Проверьте, является ли целое число степенью
// двойки. Ноль и отрицательные числа степенями двойки не считаются.
// Сложность: easy. Ожидаемая асимптотика: O(1) по времени, O(1) по памяти
package main_test

import "testing"

func isPowerOfTwo(x int) bool {
	return false
}

func TestIsPowerOfTwo(t *testing.T) {
	cases := map[int]bool{1: true, 16: true, 3: false, 0: false, -16: false, 1 << 62: true, 6: false}
	for in, want := range cases {
		if got := isPowerOfTwo(in); got != want {
			t.Errorf("isPowerOfTwo(%d) = %v, want %v", in, got, want)
		}
	}
}
