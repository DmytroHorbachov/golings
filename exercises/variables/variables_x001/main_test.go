// variables_x001: Нулевое значение
// Make the tests pass!
// I AM NOT DONE
//
// Функция должна вернуть стартовый баланс нового счёта — 100.
// Переменная объявлена через var без инициализации и получает нулевое значение.
// Тренирует: объявление переменной с начальным значением.
// Сложность: easy
package main_test

import "testing"

func startBalance() int {
	var balance int
	return balance
}

func TestStartBalance(t *testing.T) {
	if got := startBalance(); got != 100 {
		t.Errorf("startBalance() = %d, want 100", got)
	}
}
