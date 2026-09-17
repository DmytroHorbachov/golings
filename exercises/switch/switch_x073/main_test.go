// switch_x073: Несколько типов в одном case
// Make the tests pass!
// I AM NOT DONE
//
// double удваивает число типа int или float64 и возвращает его как float64.
// Код не компилируется: в ветке с двумя типами v остаётся interface{}.
// Тренирует: в case с несколькими типами переменная имеет тип выражения switch.
// Сложность: hard
package main_test

import "testing"

func double(x interface{}) float64 {
	switch v := x.(type) {
	case int, float64:
		return float64(v) * 2
	}
	return 0
}

func TestDouble(t *testing.T) {
	if got := double(21); got != 42 {
		t.Errorf("double(21) = %v", got)
	}
	if got := double(1.25); got != 2.5 {
		t.Errorf("double(1.25) = %v", got)
	}
	if got := double("x"); got != 0 {
		t.Errorf("double(x) = %v", got)
	}
}
