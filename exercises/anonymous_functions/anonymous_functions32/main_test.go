// anonymous_functions32
// Make the tests pass!

// I AM NOT DONE
//
// cleanup должен выполниться при выходе из функции. Код не компилируется.
// Тренирует: выражение в defer должно быть вызовом функции.
// Сложность: hard
package main_test

import "testing"

func work(log *[]string) {
	defer func() {
		*log = append(*log, "cleanup")
	}
	*log = append(*log, "work")
}

func TestWork(t *testing.T) {
	var log []string
	work(&log)
	if len(log) != 2 || log[1] != "cleanup" {
		t.Errorf("log = %v", log)
	}
}
