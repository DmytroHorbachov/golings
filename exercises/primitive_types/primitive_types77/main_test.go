// primitive_types77
// Make the tests pass!

// I AM NOT DONE
//
// total суммирует цены в рублях с копейками и возвращает сумму в копейках.
// Сумма десяти цен по 0.10 не равна 100 копейкам.
// Тренирует: накопление ошибки float64 в денежных расчётах.
// Сложность: hard
package main_test

import (
	"math"
	"testing"
)

func total(prices []float64) int64 {
	var sum float64
	for _, p := range prices {
		sum += p
	}
	return int64(sum * 100)
}

func TestTotal(t *testing.T) {
	_ = math.Round
	ten := make([]float64, 10)
	for i := range ten {
		ten[i] = 0.10
	}
	if got := total(ten); got != 100 {
		t.Errorf("total(10 x 0.10) = %d, want 100", got)
	}
	if got := total([]float64{19.99, 0.01, 4.35}); got != 2435 {
		t.Errorf("total = %d, want 2435", got)
	}
}
