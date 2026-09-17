// arrays_x072: Срез делит память с массивом
// Make the tests pass!
// I AM NOT DONE
//
// snapshot должна вернуть независимую копию показаний, которую не затронут
// последующие измерения. Сейчас «снимок» меняется вместе с массивом.
// Тренирует: срез массива arr[:] ссылается на тот же массив.
// Сложность: hard
package main_test

import "testing"

type Sensor struct {
	readings [4]int
}

func (s *Sensor) snapshot() []int {
	return s.readings[:]
}

func TestSnapshot(t *testing.T) {
	s := &Sensor{readings: [4]int{1, 2, 3, 4}}
	snap := s.snapshot()
	s.readings[0] = 100
	if snap[0] != 1 {
		t.Errorf("snapshot changed: %v", snap)
	}
}
