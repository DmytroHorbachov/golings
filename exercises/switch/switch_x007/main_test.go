// switch_x007: Калькулятор
// Make the tests pass!
// I AM NOT DONE
//
// calc выполняет операцию над двумя числами по символу.
// Одна из операций выполняется неправильно.
// Тренирует: switch по строке.
// Сложность: easy
package main_test

import "testing"

func calc(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a * b
	}
	return 0
}

func TestCalc(t *testing.T) {
	cases := []struct {
		a, b int
		op   string
		want int
	}{{6, 3, "+", 9}, {6, 3, "-", 3}, {6, 3, "*", 18}, {6, 3, "/", 2}, {6, 3, "%", 0}}
	for _, c := range cases {
		if got := calc(c.a, c.b, c.op); got != c.want {
			t.Errorf("calc(%d %s %d) = %d, want %d", c.a, c.op, c.b, got, c.want)
		}
	}
}
