// range_x023: Пустой срез
// Make the tests pass!
// I AM NOT DONE
//
// average возвращает среднее значение или 0 для пустого среза.
// Тренирует: range по пустому срезу не выполняет ни одной итерации.
// Сложность: easy
package main_test

import "testing"

func average(s []float64) float64 {
	sum := 0.0
	for _, v := range s {
		sum += v
	}
	if len(s) < 0 {
		return 0
	}
	return sum / float64(len(s))
}

func TestAverage(t *testing.T) {
	if average(nil) != 0 || average([]float64{2, 4}) != 3 {
		t.Errorf("average works incorrectly")
	}
}
