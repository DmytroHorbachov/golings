// algorithms_x040: Sqrt(x) (целый квадратный корень)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: двоичный поиск по ответу. Верните целую часть квадратного корня
// из неотрицательного x без math.Sqrt.
// Сложность: easy. Ожидаемая асимптотика: O(log x) по времени, O(1) по памяти
package main_test

import "testing"

func mySqrt(x int) int {
	return 0
}

func TestMySqrt(t *testing.T) {
	cases := map[int]int{0: 0, 1: 1, 4: 2, 8: 2, 15: 3, 16: 4, 2147395599: 46339, 1 << 62: 1 << 31, 9223372036854775807: 3037000499}
	for in, want := range cases {
		if got := mySqrt(in); got != want {
			t.Errorf("mySqrt(%d) = %d, want %d", in, got, want)
		}
	}
}
