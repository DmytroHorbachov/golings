// switch_x072: Условия при switch с тегом
// Make the tests pass!
// I AM NOT DONE
//
// classify возвращает "neg", "zero" или "pos". Код не компилируется:
// в case стоят логические выражения, а switch идёт по числу.
// Тренирует: switch x сравнивает x с каждым case; для условий нужен switch без тега.
// Сложность: hard
package main_test

import "testing"

func classify(n int) string {
	switch n {
	case n < 0:
		return "neg"
	case n == 0:
		return "zero"
	default:
		return "pos"
	}
}

func TestClassify(t *testing.T) {
	cases := map[int]string{-3: "neg", 0: "zero", 8: "pos"}
	for in, want := range cases {
		if got := classify(in); got != want {
			t.Errorf("classify(%d) = %s, want %s", in, got, want)
		}
	}
}
