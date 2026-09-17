// structs_x083: NaN в поле
// Make the tests pass!
// I AM NOT DONE
//
// Структуры Measurement с «пустым» значением NaN должны считаться равными.
// Тренирует: == для структур сравнивает float-поля по IEEE (NaN != NaN).
// Сложность: hard
package main_test

import (
	"math"
	"testing"
)

type Measurement struct {
	Sensor string
	Value  float64
}

func (m Measurement) Equal(o Measurement) bool {
	return m == o
}

func TestMeasurementEqual(t *testing.T) {
	nan := math.NaN()
	if !(Measurement{"t1", nan}).Equal(Measurement{"t1", nan}) {
		t.Errorf("NaN measurements should be equal")
	}
	if (Measurement{"t1", 1}).Equal(Measurement{"t2", 1}) {
		t.Errorf("different sensors should differ")
	}
}
