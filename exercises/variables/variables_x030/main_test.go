// variables_x030: Проценты на целых
// Make the tests pass!
// I AM NOT DONE
//
// Функция должна вернуть долю выполненных задач в процентах (0..100).
// Для 1 из 3 ожидается 33, а получается 0.
// Тренирует: порядок операций при целочисленном делении.
// Сложность: hard
package main_test

import "testing"

func percentDone(done, total int) int {
	if total == 0 {
		return 0
	}
	return done / total * 100
}

func TestPercentDone(t *testing.T) {
	cases := []struct{ done, total, want int }{
		{1, 3, 33}, {2, 3, 66}, {3, 3, 100}, {0, 5, 0}, {1, 0, 0},
	}
	for _, c := range cases {
		if got := percentDone(c.done, c.total); got != c.want {
			t.Errorf("percentDone(%d, %d) = %d, want %d", c.done, c.total, got, c.want)
		}
	}
}
