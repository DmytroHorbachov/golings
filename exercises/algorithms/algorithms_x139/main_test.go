// algorithms_x139: Palindrome Number (число-палиндром)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: работа с цифрами числа. Проверьте, читается ли целое число
// одинаково слева направо и справа налево, не преобразуя его в строку.
// Сложность: easy. Ожидаемая асимптотика: O(log n) по времени, O(1) по памяти
package main_test

import "testing"

func isPalindromeNumber(x int) bool {
	return false
}

func TestIsPalindromeNumber(t *testing.T) {
	cases := map[int]bool{121: true, -121: false, 10: false, 0: true, 1221: true, 12321: true, 1000021: false}
	for in, want := range cases {
		if got := isPalindromeNumber(in); got != want {
			t.Errorf("isPalindromeNumber(%d) = %v, want %v", in, got, want)
		}
	}
}
