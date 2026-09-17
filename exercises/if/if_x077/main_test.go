// if_x077: Нечётные отрицательные
// Make the tests pass!
// I AM NOT DONE
//
// isOdd должна корректно работать и для отрицательных чисел.
// Для -3 сейчас возвращается false.
// Тренирует: знак остатка от деления в Go.
// Сложность: hard
package main_test

import "testing"

func isOdd(n int) bool {
	if n%2 == 1 {
		return true
	}
	return false
}

func TestIsOdd(t *testing.T) {
	cases := map[int]bool{3: true, -3: true, 4: false, -4: false, 0: false}
	for in, want := range cases {
		if got := isOdd(in); got != want {
			t.Errorf("isOdd(%d) = %v, want %v", in, got, want)
		}
	}
}
