// anonymous_functions_x075: Литерал над копией получателя
// Make the tests pass!
// I AM NOT DONE
//
// Account.Depositor возвращает литерал пополнения. Пополнения не видны на счёте:
// метод со значимым получателем, и литерал захватил копию.
// Тренирует: замыкание захватывает переменную получателя.
// Сложность: hard
package main_test

import "testing"

type Account struct{ Balance int }

func (a Account) Depositor() func(int) {
	return func(n int) { a.Balance += n }
}

func TestDepositor(t *testing.T) {
	acc := &Account{}
	dep := acc.Depositor()
	dep(10)
	dep(5)
	if acc.Balance != 15 {
		t.Errorf("Balance = %d, want 15", acc.Balance)
	}
}
