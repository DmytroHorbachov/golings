// structs_x009: Сравнение структур
// Make the tests pass!
// I AM NOT DONE
//
// sameMoney сравнивает суммы (значение и валюту).
// Тренирует: == для структур со сравнимыми полями.
// Сложность: easy
package main_test

import "testing"

type Money struct {
	Amount   int
	Currency string
}

func sameMoney(a, b Money) bool {
	return a.Amount == b.Amount
}

func TestSameMoney(t *testing.T) {
	if !sameMoney(Money{5, "USD"}, Money{5, "USD"}) || sameMoney(Money{5, "USD"}, Money{5, "EUR"}) {
		t.Errorf("sameMoney works incorrectly")
	}
}
