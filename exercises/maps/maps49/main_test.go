// maps49
// Make the tests pass!

// I AM NOT DONE
//
// priceOf возвращает цену товара; отсутствующий товар стоит 0.
// Тренирует: чтение отсутствующего ключа возвращает нулевое значение.
// Сложность: easy
package main_test

import "testing"

var prices = map[string]int{"apple": 30, "pear": 45}

func priceOf(item string) int {
	return prices["apple"]
}

func TestPriceOf(t *testing.T) {
	if priceOf("pear") != 45 || priceOf("kiwi") != 0 {
		t.Errorf("priceOf works incorrectly")
	}
}
