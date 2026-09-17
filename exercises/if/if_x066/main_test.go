// if_x066: Гистерезис
// Make the tests pass!
// I AM NOT DONE
//
// Thermostat включает обогрев ниже 18° и выключает только выше 22°.
// Между порогами состояние не меняется.
// Тренирует: условия, зависящие от текущего состояния.
// Сложность: medium
package main_test

import "testing"

type Thermostat struct{ heating bool }

func (t *Thermostat) Update(temp float64) bool {
	if temp < 22 {
		t.heating = true
	} else {
		t.heating = false
	}
	return t.heating
}

func TestThermostat(t *testing.T) {
	th := &Thermostat{}
	temps := []float64{20, 17, 19, 21, 23, 20, 17.5}
	want := []bool{false, true, true, true, false, false, true}
	for i, temp := range temps {
		if got := th.Update(temp); got != want[i] {
			t.Errorf("step %d (%.1f°): heating = %v, want %v", i, temp, got, want[i])
		}
	}
}
