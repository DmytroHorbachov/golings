// variables80
// Make the tests pass!

// I AM NOT DONE
//
// Функция должна вернуть среднее двух целых чисел как float64.
// Сейчас деление выполняется над целыми, и дробная часть теряется.
// Тренирует: явное преобразование типов перед арифметикой.
// Сложность: easy
package main_test

import "testing"

func average(a, b int) float64 {
	avg := float64((a + b) / 2)
	return avg
}

func TestAverage(t *testing.T) {
	cases := []struct {
		a, b int
		want float64
	}{{1, 2, 1.5}, {4, 4, 4}, {-3, 0, -1.5}}
	for _, c := range cases {
		if got := average(c.a, c.b); got != c.want {
			t.Errorf("average(%d, %d) = %v, want %v", c.a, c.b, got, c.want)
		}
	}
}
