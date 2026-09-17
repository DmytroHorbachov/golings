// generics71
// Make the tests pass!

// I AM NOT DONE
//
// Box[T] должна уметь преобразовывать значение в другой тип. Код не
// компилируется: у методов не может быть собственных параметров типа.
// Тренирует: ограничение языка на методы.
// Сложность: hard
package main_test

import (
	"strconv"
	"testing"
)

type Box[T any] struct{ V T }

func (b Box[T]) Map[U any](f func(T) U) Box[U] {
	return Box[U]{f(b.V)}
}

func describe() string {
	return Box[int]{42}.Map(strconv.Itoa).V
}

func TestDescribe(t *testing.T) {
	if describe() != "42" {
		t.Errorf("describe = %q", describe())
	}
}
