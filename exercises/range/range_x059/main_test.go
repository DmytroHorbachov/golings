// range_x059: Баланс по операциям
// Make the tests pass!
// I AM NOT DONE
//
// balance применяет операции: "in" увеличивает баланс, "out" уменьшает,
// операции с неизвестным типом пропускаются.
// Тренирует: range со switch внутри.
// Сложность: medium
package main_test

import "testing"

type Op struct {
	Kind   string
	Amount int
}

func balance(ops []Op) int {
	b := 0
	for _, op := range ops {
		switch op.Kind {
		case "in", "out":
			b += op.Amount
		default:
			b -= op.Amount
		}
	}
	return b
}

func TestBalance(t *testing.T) {
	ops := []Op{{"in", 100}, {"out", 30}, {"fee", 5}, {"in", 10}}
	if got := balance(ops); got != 80 {
		t.Errorf("balance = %d, want 80", got)
	}
}
