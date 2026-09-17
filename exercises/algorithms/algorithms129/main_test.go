// algorithms129
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: жадность. Стакан стоит 5; покупатели платят купюрами 5, 10 или 20.
// Верните true, если каждому удалось дать сдачу (касса пуста в начале).
// Сложность: easy. Ожидаемая асимптотика: O(n) по времени, O(1) по памяти
package main_test

import "testing"

func lemonadeChange(bills []int) bool {
	return false
}

func TestLemonadeChange(t *testing.T) {
	cases := []struct {
		bills []int
		want  bool
	}{
		{[]int{5, 5, 5, 10, 20}, true},
		{[]int{5, 5, 10, 10, 20}, false},
		{[]int{10}, false},
		{nil, true},
		{[]int{5, 5, 5, 5, 5, 5, 20, 20}, true},
		{[]int{5, 5, 5, 5, 20, 20}, false},
	}
	for _, c := range cases {
		if got := lemonadeChange(c.bills); got != c.want {
			t.Errorf("lemonadeChange(%v) = %v, want %v", c.bills, got, c.want)
		}
	}
}
