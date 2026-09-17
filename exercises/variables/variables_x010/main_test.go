// variables_x010: Составное присваивание
// Make the tests pass!
// I AM NOT DONE
//
// Функция должна вернуть 2 в степени n, удваивая результат в цикле.
// Тренирует: операторы составного присваивания (+=, *=, ...).
// Сложность: easy
package main_test

import "testing"

func powerOfTwo(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result += 2
	}
	return result
}

func TestPowerOfTwo(t *testing.T) {
	cases := map[int]int{0: 1, 1: 2, 5: 32, 10: 1024}
	for n, want := range cases {
		if got := powerOfTwo(n); got != want {
			t.Errorf("powerOfTwo(%d) = %d, want %d", n, got, want)
		}
	}
}
