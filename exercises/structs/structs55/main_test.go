// structs55
// Make the tests pass!

// I AM NOT DONE
//
// sameOrder сравнивает два заказа. Код не компилируется: у структуры есть поле-срез.
// Тренирует: == доступно только для структур со сравнимыми полями.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

type Order struct {
	ID    int
	Items []string
}

func sameOrder(a, b Order) bool {
	return a == b
}

func TestSameOrder(t *testing.T) {
	_ = reflect.DeepEqual
	if !sameOrder(Order{1, []string{"a"}}, Order{1, []string{"a"}}) || sameOrder(Order{1, []string{"a"}}, Order{1, []string{"b"}}) {
		t.Errorf("sameOrder works incorrectly")
	}
}
