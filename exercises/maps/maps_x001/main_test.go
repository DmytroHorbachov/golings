// maps_x001: Литерал map
// Make the tests pass!
// I AM NOT DONE
//
// capitals должна вернуть столицы стран; одна из них указана неверно.
// Тренирует: литерал map.
// Сложность: easy
package main_test

import "testing"

func capitals() map[string]string {
	return map[string]string{
		"France": "Paris",
		"Italy":  "Milan",
	}
}

func TestCapitals(t *testing.T) {
	c := capitals()
	if c["France"] != "Paris" || c["Italy"] != "Rome" {
		t.Errorf("capitals = %v", c)
	}
}
