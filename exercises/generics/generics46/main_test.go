// generics46
// Make the tests pass!

// I AM NOT DONE
//
// Scale умножает значение на 1000. Код не компилируется: в ограничении есть int8,
// в который 1000 не помещается.
// Тренирует: константы в обобщённом коде должны подходить всем типам из набора.
// Сложность: hard
package main_test

import "testing"

type Num interface {
	~int8 | ~int | ~int64
}

func Scale[T Num](v T) T {
	return v * 1000
}

func TestScale(t *testing.T) {
	if Scale(3) != 3000 || Scale(int64(2)) != 2000 {
		t.Errorf("Scale works incorrectly")
	}
}
