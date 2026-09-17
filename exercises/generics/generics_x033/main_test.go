// generics_x033: Недостающий тип в ограничении
// Make the tests pass!
// I AM NOT DONE
//
// Total вызывается для среза int32, но ограничение этот тип не содержит.
// Тренирует: объединение типов в ограничении.
// Сложность: easy
package main_test

import "testing"

type Integer interface {
	~int | ~int64
}

func Total[T Integer](s []T) T {
	var t T
	for _, v := range s {
		t += v
	}
	return t
}

func TestTotal(t *testing.T) {
	if Total([]int32{1, 2, 3}) != 6 {
		t.Errorf("Total = %d", Total([]int32{1, 2, 3}))
	}
}
