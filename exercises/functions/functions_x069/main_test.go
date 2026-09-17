// functions_x069: append в функции
// Make the tests pass!
// I AM NOT DONE
//
// addItem должна добавить элемент в корзину вызывающего кода.
// После вызова длина корзины у вызывающего не меняется.
// Тренирует: срез передаётся по значению (заголовок копируется).
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

func addItem(cart []string, item string) {
	cart = append(cart, item)
}

func fillCart() []string {
	cart := make([]string, 0, 10)
	addItem(cart, "apple")
	addItem(cart, "milk")
	return cart
}

func TestFillCart(t *testing.T) {
	if got := fillCart(); !reflect.DeepEqual(got, []string{"apple", "milk"}) {
		t.Errorf("fillCart() = %v, want [apple milk]", got)
	}
}
