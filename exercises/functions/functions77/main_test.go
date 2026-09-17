// functions77
// Make the tests pass!

// I AM NOT DONE
//
// fib должна использовать кэш, переданный в параметре, и работать быстро для n=80.
// Сейчас кэш не заполняется и не читается.
// Тренирует: передачу map в функцию и рекурсию с мемоизацией.
// Сложность: medium
package main_test

import "testing"

func fib(n int, memo map[int]int) int {
	if n < 2 {
		return n
	}
	return fib(n-1, memo) + fib(n-2, memo)
}

func TestFibMemo(t *testing.T) {
	memo := map[int]int{}
	if got := fib(80, memo); got != 23416728348467685 {
		t.Errorf("fib(80) = %d", got)
	}
	if len(memo) != 79 {
		t.Errorf("memo should contain 79 entries, got %d", len(memo))
	}
}
