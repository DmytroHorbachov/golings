// primitive_types_x004: Два знака после запятой
// Make the tests pass!
// I AM NOT DONE
//
// price должна форматировать число с двумя знаками после точки.
// Тренирует: точность в глаголе %f.
// Сложность: easy
package main_test

import (
	"fmt"
	"testing"
)

func price(p float64) string {
	return fmt.Sprintf("%.1f", p)
}

func TestPrice(t *testing.T) {
	cases := map[float64]string{3.14159: "3.14", 2: "2.00", 0.005: "0.01"}
	for in, want := range cases {
		if got := price(in); got != want {
			t.Errorf("price(%v) = %s, want %s", in, got, want)
		}
	}
}
