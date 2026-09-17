// functions_x014: Сообщение об ошибке
// Make the tests pass!
// I AM NOT DONE
//
// Функция withdraw должна вернуть ошибку с текстом "insufficient funds".
// Тренирует: создание ошибок через errors.New.
// Сложность: easy
package main_test

import (
	"errors"
	"testing"
)

func withdraw(balance, amount int) (int, error) {
	if amount > balance {
		return balance, errors.New("not enough money")
	}
	return balance - amount, nil
}

func TestWithdraw(t *testing.T) {
	if b, err := withdraw(100, 30); err != nil || b != 70 {
		t.Errorf("withdraw(100, 30) = %d, %v", b, err)
	}
	_, err := withdraw(10, 30)
	if err == nil || err.Error() != "insufficient funds" {
		t.Errorf("withdraw(10, 30) error = %v, want insufficient funds", err)
	}
}
