// generics20
// Make the tests pass!

// I AM NOT DONE
//
// Total суммирует значения Cents и возвращает T. Результат нельзя сложить с int
// без преобразования.
// Тренирует: результат сохраняет именованный тип аргумента.
// Сложность: hard
package main_test

import "testing"

type Cents int

type Integer interface{ ~int | ~int64 }

func Total[T Integer](vals ...T) T {
	var s T
	for _, v := range vals {
		s += v
	}
	return s
}

func withFee(fee int) int {
	sum := Total(Cents(100), Cents(250))
	return sum + fee
}

func TestWithFee(t *testing.T) {
	if withFee(50) != 400 {
		t.Errorf("withFee = %d", withFee(50))
	}
}
