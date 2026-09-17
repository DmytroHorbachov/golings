// structs80
// Make the tests pass!

// I AM NOT DONE
//
// fmt не использует String() при печати значения, если метод объявлен на указателе.
// Тренирует: интерфейс fmt.Stringer и наборы методов.
// Сложность: hard
package main_test

import (
	"fmt"
	"testing"
)

type Card struct {
	Rank string
	Suit string
}

func (c *Card) String() string { return c.Rank + c.Suit }

func TestCardString(t *testing.T) {
	hand := []Card{{"A", "♠"}, {"10", "♥"}}
	if got := fmt.Sprint(hand); got != "[A♠ 10♥]" {
		t.Errorf("Sprint = %q", got)
	}
}
