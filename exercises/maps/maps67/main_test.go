// maps67
// Make the tests pass!

// I AM NOT DONE
//
// cityPopulation читает население города из map стран.
// Тренирует: двойную индексацию map[string]map[string]int.
// Сложность: easy
package main_test

import "testing"

var population = map[string]map[string]int{
	"RU": {"Moscow": 13, "Kazan": 1},
	"DE": {"Berlin": 4},
}

func cityPopulation(country, city string) int {
	return population[city][country]
}

func TestCityPopulation(t *testing.T) {
	if cityPopulation("RU", "Kazan") != 1 || cityPopulation("DE", "Berlin") != 4 || cityPopulation("FR", "Paris") != 0 {
		t.Errorf("cityPopulation works incorrectly")
	}
}
