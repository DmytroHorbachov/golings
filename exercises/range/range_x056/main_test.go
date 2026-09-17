// range_x056: Средневзвешенное
// Make the tests pass!
// I AM NOT DONE
//
// weightedAvg считает средневзвешенную оценку по парам (оценка, вес).
// Тренирует: range по срезу пар и два аккумулятора.
// Сложность: medium
package main_test

import "testing"

type Grade struct{ Score, Weight float64 }

func weightedAvg(gs []Grade) float64 {
	var sum, weights float64
	for _, g := range gs {
		sum += g.Score
		weights++
	}
	return sum / weights
}

func TestWeightedAvg(t *testing.T) {
	if got := weightedAvg([]Grade{{5, 3}, {2, 1}}); got != 4.25 {
		t.Errorf("weightedAvg = %v, want 4.25", got)
	}
	if got := weightedAvg(nil); got != 0 {
		t.Errorf("weightedAvg(nil) = %v", got)
	}
}
