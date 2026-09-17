// arrays45
// Make the tests pass!

// I AM NOT DONE
//
// samePoint сравнивает точки [2]float64 и должна считать равными
// точки с «пропущенными» (NaN) координатами на тех же позициях.
// Тренирует: == для массивов сравнивает элементы, а NaN != NaN.
// Сложность: hard
package main_test

import (
	"math"
	"testing"
)

func samePoint(a, b [2]float64) bool {
	return a == b
}

func TestSamePoint(t *testing.T) {
	nan := math.NaN()
	if !samePoint([2]float64{1, nan}, [2]float64{1, nan}) {
		t.Errorf("points with NaN in the same place should match")
	}
	if samePoint([2]float64{1, nan}, [2]float64{nan, 1}) || samePoint([2]float64{1, 2}, [2]float64{1, 3}) {
		t.Errorf("different points should not match")
	}
}
