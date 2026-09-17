// range_x082: Метод на копии
// Make the tests pass!
// I AM NOT DONE
//
// depositAll пополняет все счета, вызывая метод с указателем-получателем.
// Метод вызывается на переменной цикла, и балансы в срезе не меняются.
// Тренирует: v в range — копия, и pointer-метод меняет эту копию.
// Сложность: hard
package main_test

import "testing"

type Account struct{ Balance int }

func (a *Account) Deposit(n int) { a.Balance += n }

func depositAll(accounts []Account, n int) {
	for _, acc := range accounts {
		acc.Deposit(n)
	}
}

func TestDepositAll(t *testing.T) {
	accs := []Account{{10}, {20}}
	depositAll(accs, 5)
	if accs[0].Balance != 15 || accs[1].Balance != 25 {
		t.Errorf("balances = %v", accs)
	}
}
