// variables_x060: Типизированные единицы
// Make the tests pass!
// I AM NOT DONE
//
// Функции переводят километры в мили и обратно.
// Коэффициент задан неверно, и один перевод перепутан.
// Тренирует: именованные типы, константы и преобразования.
// Сложность: medium
package main_test

import (
	"math"
	"testing"
)

type Km float64
type Mile float64

const kmPerMile = 1.6

func toMiles(k Km) Mile {
	return Mile(k * kmPerMile)
}

func toKm(m Mile) Km {
	return Km(m * kmPerMile)
}

func TestConversions(t *testing.T) {
	if got := toKm(1); math.Abs(float64(got)-1.609344) > 1e-9 {
		t.Errorf("toKm(1) = %v, want 1.609344", got)
	}
	if got := toMiles(1.609344); math.Abs(float64(got)-1) > 1e-9 {
		t.Errorf("toMiles(1.609344) = %v, want 1", got)
	}
}
