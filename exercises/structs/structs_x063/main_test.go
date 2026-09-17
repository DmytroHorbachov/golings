// structs_x063: Рациональные числа
// Make the tests pass!
// I AM NOT DONE
//
// NewRat создаёт дробь в несократимом виде со знаком в числителе;
// Add складывает дроби.
// Тренирует: инварианты структуры, поддерживаемые конструктором.
// Сложность: medium
package main_test

import "testing"

type Rat struct{ Num, Den int }

func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func NewRat(n, d int) Rat {
	return Rat{n, d}
}

func (r Rat) Add(o Rat) Rat {
	return NewRat(r.Num*o.Den+o.Num*r.Den, r.Den*o.Den)
}

func TestRat(t *testing.T) {
	if NewRat(2, -4) != (Rat{-1, 2}) {
		t.Errorf("NewRat(2, -4) = %v", NewRat(2, -4))
	}
	if got := NewRat(1, 6).Add(NewRat(1, 3)); got != (Rat{1, 2}) {
		t.Errorf("1/6 + 1/3 = %v", got)
	}
}
