// algorithms113
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: решето Эратосфена. Посчитайте количество простых чисел,
// строго меньших n. Решение должно быть быстрее перебора делителей.
// Сложность: medium. Ожидаемая асимптотика: O(n·log log n) по времени, O(n) по памяти
package main_test

import "testing"

func countPrimes(n int) int {
	return 0
}

func TestCountPrimes(t *testing.T) {
	cases := map[int]int{10: 4, 0: 0, 1: 0, 2: 0, 3: 1, 100: 25, 1000000: 78498}
	for in, want := range cases {
		if got := countPrimes(in); got != want {
			t.Errorf("countPrimes(%d) = %d, want %d", in, got, want)
		}
	}
}
