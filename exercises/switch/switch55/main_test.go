// switch55
// Make the tests pass!

// I AM NOT DONE
//
// points переводит буквенную оценку в баллы: A=5, B=4, C=3, остальное 0.
// Тренирует: switch по байту.
// Сложность: easy
package main_test

import "testing"

func points(g byte) int {
	switch g {
	case 'A':
		return 5
	case 'B':
		return 3
	case 'C':
		return 3
	}
	return 0
}

func TestPoints(t *testing.T) {
	cases := map[byte]int{'A': 5, 'B': 4, 'C': 3, 'F': 0}
	for in, want := range cases {
		if got := points(in); got != want {
			t.Errorf("points(%c) = %d, want %d", in, got, want)
		}
	}
}
