// generics57
// Make the tests pass!

// I AM NOT DONE
//
// Double удваивает значения, в том числе именованного типа Celsius.
// Код не компилируется: ограничение принимает только сам float64.
// Тренирует: ~T включает все типы с базовым типом T.
// Сложность: easy
package main_test

import "testing"

type Celsius float64

type Float interface {
	float64
}

func Double[T Float](v T) T { return v * 2 }

func TestDouble(t *testing.T) {
	if Double(Celsius(21.5)) != 43 {
		t.Errorf("Double = %v", Double(Celsius(21.5)))
	}
}
