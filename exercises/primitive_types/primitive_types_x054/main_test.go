// primitive_types_x054: Сложение с проверкой
// Make the tests pass!
// I AM NOT DONE
//
// addChecked складывает int64 и сообщает о переполнении.
// Тренирует: обнаружение переполнения через знаки операндов и результата.
// Сложность: medium
package main_test

import (
	"math"
	"testing"
)

func addChecked(a, b int64) (int64, bool) {
	s := a + b
	if s < a || s < b {
		return s, false
	}
	return s, true
}

func TestAddChecked(t *testing.T) {
	if s, ok := addChecked(2, 3); !ok || s != 5 {
		t.Errorf("addChecked(2, 3) = %d, %v", s, ok)
	}
	if s, ok := addChecked(5, -3); !ok || s != 2 {
		t.Errorf("addChecked(5, -3) = %d, %v", s, ok)
	}
	if _, ok := addChecked(math.MaxInt64, 1); ok {
		t.Errorf("MaxInt64+1 should overflow")
	}
	if _, ok := addChecked(math.MinInt64, -1); ok {
		t.Errorf("MinInt64-1 should overflow")
	}
}
