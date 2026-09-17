// if_x083: Законы де Моргана
// Make the tests pass!
// I AM NOT DONE
//
// needsReview: заказ НЕ требует проверки, только если он оплачен И сумма
// меньше 10000. Во всех остальных случаях проверка нужна.
// Тренирует: корректное отрицание составного условия.
// Сложность: hard
package main_test

import "testing"

func needsReview(paid bool, total int) bool {
	if !paid && total >= 10000 {
		return true
	}
	return false
}

func TestNeedsReview(t *testing.T) {
	cases := []struct {
		paid  bool
		total int
		want  bool
	}{{true, 500, false}, {false, 500, true}, {true, 20000, true}, {false, 20000, true}}
	for _, c := range cases {
		if got := needsReview(c.paid, c.total); got != c.want {
			t.Errorf("needsReview(%v, %d) = %v, want %v", c.paid, c.total, got, c.want)
		}
	}
}
