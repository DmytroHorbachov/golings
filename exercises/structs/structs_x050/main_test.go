// structs_x050: Склад
// Make the tests pass!
// I AM NOT DONE
//
// Inventory.Remove уменьшает количество и удаляет товар при нуле;
// нельзя убрать больше, чем есть.
// Тренирует: методы с проверками над полем-map.
// Сложность: medium
package main_test

import (
	"errors"
	"testing"
)

type Inventory struct{ stock map[string]int }

func (inv *Inventory) Remove(item string, n int) error {
	inv.stock[item] -= n
	return nil
}

func TestInventory(t *testing.T) {
	_ = errors.New
	inv := &Inventory{stock: map[string]int{"pen": 3}}
	if err := inv.Remove("pen", 5); err == nil {
		t.Errorf("removing too much should fail")
	}
	inv.Remove("pen", 1)
	inv.Remove("pen", 2)
	if _, ok := inv.stock["pen"]; ok {
		t.Errorf("pen should be removed: %v", inv.stock)
	}
}
