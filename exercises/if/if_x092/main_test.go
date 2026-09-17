// if_x092: float32 против float64
// Make the tests pass!
// I AM NOT DONE
//
// isDefaultRatio проверяет, что коэффициент (float32) равен 0.1.
// Сравнение после преобразования в float64 никогда не срабатывает.
// Тренирует: 0.1 по-разному округляется в float32 и float64.
// Сложность: hard
package main_test

import "testing"

func isDefaultRatio(r float32) bool {
	if float64(r) == 0.1 {
		return true
	}
	return false
}

func TestIsDefaultRatio(t *testing.T) {
	var r float32 = 0.1
	if !isDefaultRatio(r) {
		t.Errorf("isDefaultRatio(0.1) = false, want true")
	}
	if isDefaultRatio(0.2) {
		t.Errorf("isDefaultRatio(0.2) = true, want false")
	}
}
