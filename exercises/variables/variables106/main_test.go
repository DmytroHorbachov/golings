// variables106
// Make the tests pass!

// I AM NOT DONE
//
// Функция tax должна вернуть налог 20% от цены в копейках (с усечением).
// Код не компилируется: дробная константа не может участвовать в int-выражении.
// Тренирует: нетипизированные константы и явные преобразования.
// Сложность: easy
package main_test

import "testing"

const taxRate = 0.2

func tax(price int) int {
	return price * taxRate
}

func TestTax(t *testing.T) {
	cases := map[int]int{100: 20, 999: 199, 0: 0}
	for in, want := range cases {
		if got := tax(in); got != want {
			t.Errorf("tax(%d) = %d, want %d", in, got, want)
		}
	}
}
