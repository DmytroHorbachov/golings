// switch_x048: День недели по смещению
// Make the tests pass!
// I AM NOT DONE
//
// weekday возвращает название дня для номера n, где 0 — понедельник.
// Номер может быть любым целым, в том числе отрицательным.
// Тренирует: нормализацию значения перед switch.
// Сложность: medium
package main_test

import "testing"

func weekday(n int) string {
	switch n {
	case 0:
		return "Mon"
	case 1:
		return "Tue"
	case 2:
		return "Wed"
	case 3:
		return "Thu"
	case 4:
		return "Fri"
	case 5:
		return "Sat"
	case 7:
		return "Sun"
	}
	return ""
}

func TestWeekday(t *testing.T) {
	cases := map[int]string{0: "Mon", 6: "Sun", 7: "Mon", 15: "Tue", -1: "Sun", -8: "Sun"}
	for in, want := range cases {
		if got := weekday(in); got != want {
			t.Errorf("weekday(%d) = %s, want %s", in, got, want)
		}
	}
}
