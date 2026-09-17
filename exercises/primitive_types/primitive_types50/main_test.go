// primitive_types50
// Make the tests pass!

// I AM NOT DONE
//
// area принимает стороны как int32 и должна вернуть площадь как int64.
// Для больших сторон результат переполняется ещё до преобразования.
// Тренирует: порядок преобразования типов при умножении.
// Сложность: easy
package main_test

import "testing"

func area(w, h int32) int64 {
	return int64(w * h)
}

func TestArea(t *testing.T) {
	if got := area(3, 4); got != 12 {
		t.Errorf("area(3, 4) = %d", got)
	}
	if got := area(100000, 100000); got != 10000000000 {
		t.Errorf("area(1e5, 1e5) = %d, want 10000000000", got)
	}
}
