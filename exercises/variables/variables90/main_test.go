// variables90
// Make the tests pass!

// I AM NOT DONE
//
// Функция должна вернуть номер заказа в виде строки "Order #42".
// Вместо цифр в строке оказывается странный символ.
// Тренирует: разницу между string(int) и strconv.Itoa.
// Сложность: hard
package main_test

import (
	"strconv"
	"testing"
)

func orderLabel(id int) string {
	_ = strconv.Itoa
	return "Order #" + string(rune(id))
}

func TestOrderLabel(t *testing.T) {
	if got := orderLabel(42); got != "Order #42" {
		t.Errorf("orderLabel(42) = %q, want %q", got, "Order #42")
	}
	if got := orderLabel(7); got != "Order #7" {
		t.Errorf("orderLabel(7) = %q, want %q", got, "Order #7")
	}
}
