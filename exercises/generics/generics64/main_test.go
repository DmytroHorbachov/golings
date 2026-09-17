// generics64
// Make the tests pass!

// I AM NOT DONE
//
// Negate возвращает -v, но в ограничение попали беззнаковые типы.
// Код не компилируется.
// Тренирует: -1 не представимо в uint, поэтому операции со знаком требуют знаковых типов.
// Сложность: hard
package main_test

import "testing"

type Signed interface {
	~int | ~int64 | ~uint
}

func Negate[T Signed](v T) T {
	return v * -1
}

func TestNegate(t *testing.T) {
	if Negate(5) != -5 || Negate(int64(-3)) != 3 {
		t.Errorf("Negate works incorrectly")
	}
}
