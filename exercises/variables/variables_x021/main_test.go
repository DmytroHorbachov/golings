// variables_x021: Шаг Фибоначчи
// Make the tests pass!
// I AM NOT DONE
//
// Функция должна вернуть n-е число Фибоначчи (fib(0)=0, fib(1)=1).
// Тренирует: одновременное присваивание нескольких переменных.
// Сложность: medium
package main_test

import "testing"

func fib(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a = b
		b = a + b
	}
	return a
}

func TestFib(t *testing.T) {
	want := []int{0, 1, 1, 2, 3, 5, 8, 13, 21}
	for n, w := range want {
		if got := fib(n); got != w {
			t.Errorf("fib(%d) = %d, want %d", n, got, w)
		}
	}
}
