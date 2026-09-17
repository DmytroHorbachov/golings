// switch59
// Make the tests pass!

// I AM NOT DONE
//
// permString превращает число 0–7 в строку прав вида "r-x".
// Тренирует: switch по отдельным битам.
// Сложность: medium
package main_test

import "testing"

func bit(p, mask int, ch byte) byte {
	switch p & mask {
	case 0:
		return ch
	default:
		return '-'
	}
}

func permString(p int) string {
	return string([]byte{bit(p, 1, 'r'), bit(p, 2, 'w'), bit(p, 4, 'x')})
}

func TestPermString(t *testing.T) {
	cases := map[int]string{0: "---", 7: "rwx", 5: "r-x", 6: "rw-", 1: "--x"}
	for in, want := range cases {
		if got := permString(in); got != want {
			t.Errorf("permString(%d) = %s, want %s", in, got, want)
		}
	}
}
