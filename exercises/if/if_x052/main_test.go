// if_x052: Стоимость парковки
// Make the tests pass!
// I AM NOT DONE
//
// parkingFee: первые 60 минут бесплатно, затем 100 за каждый начатый час.
// Тренирует: условия и округление вверх.
// Сложность: medium
package main_test

import "testing"

func parkingFee(minutes int) int {
	if minutes <= 60 {
		return 0
	}
	hours := minutes / 60
	return hours * 100
}

func TestParkingFee(t *testing.T) {
	cases := map[int]int{30: 0, 60: 0, 61: 100, 120: 100, 121: 200, 185: 300}
	for in, want := range cases {
		if got := parkingFee(in); got != want {
			t.Errorf("parkingFee(%d) = %d, want %d", in, got, want)
		}
	}
}
