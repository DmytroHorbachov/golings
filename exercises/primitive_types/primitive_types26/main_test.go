// primitive_types26
// Make the tests pass!

// I AM NOT DONE
//
// plain должна вывести число без экспоненты и без лишних нулей: 1e21 -> "1000000000000000000000".
// Тренирует: %v для больших float64 переходит на экспоненциальную форму.
// Сложность: hard
package main_test

import (
	"fmt"
	"strconv"
	"testing"
)

func plain(x float64) string {
	return fmt.Sprintf("%v", x)
}

func TestPlain(t *testing.T) {
	_, _ = fmt.Sprint, strconv.FormatFloat
	cases := map[float64]string{1e21: "1000000000000000000000", 0.000001: "0.000001", 2.5: "2.5", 100: "100"}
	for in, want := range cases {
		if got := plain(in); got != want {
			t.Errorf("plain(%v) = %s, want %s", in, got, want)
		}
	}
}
