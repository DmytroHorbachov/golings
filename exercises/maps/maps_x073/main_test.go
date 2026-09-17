// maps_x073: NaN как ключ
// Make the tests pass!
// I AM NOT DONE
//
// countReadings считает, сколько раз встретилось каждое показание; NaN
// («нет данных») должен учитываться под одним ключом.
// Сейчас каждое NaN создаёт новый ключ, и прочитать его нельзя.
// Тренирует: NaN != NaN, поэтому NaN-ключи в map уникальны и недостижимы.
// Сложность: hard
package main_test

import (
	"math"
	"testing"
)

func countReadings(vals []float64) (map[float64]int, int) {
	m := map[float64]int{}
	missing := 0
	for _, v := range vals {
		m[v]++
	}
	return m, missing
}

func TestCountReadings(t *testing.T) {
	nan := math.NaN()
	m, missing := countReadings([]float64{1.5, nan, 1.5, nan, nan})
	if m[1.5] != 2 || missing != 3 || len(m) != 1 {
		t.Errorf("m = %v, missing = %d", m, missing)
	}
}
