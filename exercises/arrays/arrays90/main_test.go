// arrays90
// Make the tests pass!

// I AM NOT DONE
//
// primes должна вернуть массив первых пяти простых чисел; размер
// выводится компилятором из литерала.
// Тренирует: литерал [...]T{...}.
// Сложность: easy
package main_test

import "testing"

func primes() [5]int {
	p := [...]int{2, 3, 5, 7, 9}
	return p
}

func TestPrimes(t *testing.T) {
	if got := primes(); got != [5]int{2, 3, 5, 7, 11} {
		t.Errorf("primes = %v", got)
	}
}
