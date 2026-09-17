// generics_x095: Литерал обобщённой структуры
// Make the tests pass!
// I AM NOT DONE
//
// Pair{1, "a"} не компилируется: для составных литералов аргументы типа
// не выводятся.
// Тренирует: инстанциация типа указывается явно или через функцию-конструктор.
// Сложность: hard
package main_test

import "testing"

type Pair[A, B any] struct {
	First  A
	Second B
}

func sample() Pair[int, string] {
	return Pair{1, "a"}
}

func TestSample(t *testing.T) {
	if p := sample(); p.First != 1 || p.Second != "a" {
		t.Errorf("sample = %+v", p)
	}
}
