// structs70
// Make the tests pass!

// I AM NOT DONE
//
// Account.Balance возвращает баланс в рублях из копеек.
// Тренирует: методы-геттеры и неэкспортируемые поля.
// Сложность: easy
package main_test

import "testing"

type Account struct{ cents int }

func (a Account) Balance() int {
	return a.cents * 100
}

func TestBalance(t *testing.T) {
	if got := (Account{cents: 12345}).Balance(); got != 123 {
		t.Errorf("Balance = %d", got)
	}
}
