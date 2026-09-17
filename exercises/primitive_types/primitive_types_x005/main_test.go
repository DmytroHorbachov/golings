// primitive_types_x005: Максимум uint8
// Make the tests pass!
// I AM NOT DONE
//
// maxByte должна вернуть наибольшее значение типа uint8.
// Тренирует: диапазон беззнаковых типов и константы math.
// Сложность: easy
package main_test

import (
	"math"
	"testing"
)

func maxByte() uint8 {
	return math.MaxInt8
}

func TestMaxByte(t *testing.T) {
	if got := maxByte(); got != 255 {
		t.Errorf("maxByte() = %d, want 255", got)
	}
}
