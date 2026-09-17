// if_x086: int и int64 в интерфейсе
// Make the tests pass!
// I AM NOT DONE
//
// isInteger должна вернуть true для значения, пришедшего как обычная целая константа.
// Сейчас проверка на int64 не срабатывает.
// Тренирует: нетипизированная константа в interface{} получает тип int.
// Сложность: hard
package main_test

import "testing"

func isInteger(v interface{}) bool {
	if _, ok := v.(int64); ok {
		return true
	}
	return false
}

func TestIsInteger(t *testing.T) {
	if !isInteger(42) {
		t.Errorf("isInteger(42) = false, want true")
	}
	if isInteger("42") || isInteger(4.2) {
		t.Errorf("strings and floats are not integers")
	}
}
