// variables_x089: Подсчёт битов
// Make the tests pass!
// I AM NOT DONE
//
// Функция bitInfo должна вернуть количество единичных битов и длину числа в битах.
// Тренирует: пакет math/bits и беззнаковые типы.
// Сложность: medium
package main_test

import (
	"math/bits"
	"testing"
)

func bitInfo(v uint) (ones, length int) {
	ones = bits.TrailingZeros(v)
	length = bits.LeadingZeros(v)
	return
}

func TestBitInfo(t *testing.T) {
	cases := []struct {
		v            uint
		ones, length int
	}{{0b1011, 3, 4}, {1, 1, 1}, {0, 0, 0}, {255, 8, 8}}
	for _, c := range cases {
		ones, length := bitInfo(c.v)
		if ones != c.ones || length != c.length {
			t.Errorf("bitInfo(%b) = %d, %d; want %d, %d", c.v, ones, length, c.ones, c.length)
		}
	}
}
