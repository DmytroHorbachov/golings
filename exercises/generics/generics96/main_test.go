// generics96
// Make the tests pass!

// I AM NOT DONE
//
// Обобщённую функцию сохраняют в переменную. Код не компилируется:
// её нельзя использовать без инстанциации.
// Тренирует: значение функции требует конкретных аргументов типа.
// Сложность: hard
package main_test

import "testing"

type Number interface{ ~int | ~float64 }

func Double[T Number](v T) T { return v * 2 }

func applyAll(vals []int, f func(int) int) []int {
	out := make([]int, len(vals))
	for i, v := range vals {
		out[i] = f(v)
	}
	return out
}

func doubled() []int {
	f := Double
	return applyAll([]int{1, 2}, f)
}

func TestDoubled(t *testing.T) {
	if got := doubled(); got[0] != 2 || got[1] != 4 {
		t.Errorf("doubled = %v", got)
	}
}
