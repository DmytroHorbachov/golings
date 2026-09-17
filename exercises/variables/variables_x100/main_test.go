// variables_x100: Порог в именованном типе
// Make the tests pass!
// I AM NOT DONE
//
// Функция alerts должна вернуть количество показаний выше порога Threshold.
// Порог и показания имеют тип Celsius; сейчас код не компилируется.
// Тренирует: типизированные константы и совместимость типов.
// Сложность: medium
package main_test

import "testing"

type Celsius float64

const Threshold Celsius = 37.5

func alerts(readings []float64) int {
	n := 0
	for _, r := range readings {
		if r >= Threshold {
			n++
		}
	}
	return n
}

func TestAlerts(t *testing.T) {
	if got := alerts([]Celsius{36.6, 37.5, 38.2, 39}); got != 2 {
		t.Errorf("alerts = %d, want 2", got)
	}
}
