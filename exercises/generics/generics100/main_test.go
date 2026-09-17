// generics100
// Make the tests pass!

// I AM NOT DONE
//
// SumAll складывает значения типа Cents (type Cents int).
// Код не компилируется: ограничение принимает только int.
// Тренирует: без ~ именованные типы не входят в набор типов.
// Сложность: hard
package main_test

import "testing"

type Cents int

type Integer interface {
	int | int64
}

func SumAll[T Integer](vals ...T) T {
	var s T
	for _, v := range vals {
		s += v
	}
	return s
}

func TestSumAll(t *testing.T) {
	if SumAll(Cents(150), Cents(250)) != 400 {
		t.Errorf("SumAll = %d", SumAll(Cents(150), Cents(250)))
	}
}
