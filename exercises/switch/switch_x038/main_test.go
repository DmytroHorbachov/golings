// switch_x038: Время года по месяцу
// Make the tests pass!
// I AM NOT DONE
//
// season возвращает время года для номера месяца (1–12) и "invalid" иначе.
// Тренирует: списки значений в case.
// Сложность: medium
package main_test

import "testing"

func season(m int) string {
	switch m {
	case 1, 2, 3:
		return "winter"
	case 4, 5:
		return "spring"
	case 6, 7, 8:
		return "summer"
	case 9, 10, 11:
		return "autumn"
	}
	return "invalid"
}

func TestSeason(t *testing.T) {
	cases := map[int]string{12: "winter", 2: "winter", 3: "spring", 5: "spring", 7: "summer", 10: "autumn", 13: "invalid"}
	for in, want := range cases {
		if got := season(in); got != want {
			t.Errorf("season(%d) = %s, want %s", in, got, want)
		}
	}
}
