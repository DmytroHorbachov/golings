// variables99
// Make the tests pass!

// I AM NOT DONE
//
// Функция pages должна вернуть число страниц для total элементов по perPage на странице.
// Неполная страница тоже считается.
// Тренирует: целочисленное деление с округлением вверх.
// Сложность: medium
package main_test

import "testing"

func pages(total, perPage int) int {
	return total / perPage
}

func TestPages(t *testing.T) {
	cases := []struct{ total, per, want int }{
		{25, 10, 3}, {30, 10, 3}, {0, 10, 0}, {1, 10, 1}, {5, 0, 0},
	}
	for _, c := range cases {
		if got := pages(c.total, c.per); got != c.want {
			t.Errorf("pages(%d, %d) = %d, want %d", c.total, c.per, got, c.want)
		}
	}
}
