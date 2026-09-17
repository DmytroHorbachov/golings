// switch77
// Make the tests pass!

// I AM NOT DONE
//
// reading возвращает "missing" для NaN, иначе "low" (<10) или "high".
// Ветка case math.NaN() никогда не срабатывает.
// Тренирует: NaN не равен ничему, поэтому switch по значению его не поймает.
// Сложность: hard
package main_test

import (
	"math"
	"testing"
)

func reading(v float64) string {
	switch v {
	case math.NaN():
		return "missing"
	}
	if v < 10 {
		return "low"
	}
	return "high"
}

func TestReading(t *testing.T) {
	cases := map[float64]string{3: "low", 42: "high"}
	for in, want := range cases {
		if got := reading(in); got != want {
			t.Errorf("reading(%v) = %s, want %s", in, got, want)
		}
	}
	if got := reading(math.NaN()); got != "missing" {
		t.Errorf("reading(NaN) = %s, want missing", got)
	}
}
