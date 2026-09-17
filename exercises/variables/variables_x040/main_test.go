// variables_x040: Это не степень
// Make the tests pass!
// I AM NOT DONE
//
// Функция square должна возвращать квадрат числа.
// Оператор ^ выглядит как возведение в степень, но это не так.
// Тренирует: побитовые операторы Go.
// Сложность: hard
package main_test

import "testing"

func square(x int) int {
	return x ^ 2
}

func TestSquare(t *testing.T) {
	cases := map[int]int{0: 0, 1: 1, 3: 9, -4: 16, 10: 100}
	for in, want := range cases {
		if got := square(in); got != want {
			t.Errorf("square(%d) = %d, want %d", in, got, want)
		}
	}
}
