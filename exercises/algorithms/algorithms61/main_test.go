// algorithms61
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: эйлеров путь (алгоритм Хирхольцера). По списку билетов [from, to]
// постройте маршрут, начинающийся в "JFK" и использующий все билеты.
// При выборе берите лексикографически меньший аэропорт.
// Сложность: hard. Ожидаемая асимптотика: O(E·log E) по времени, O(E) по памяти
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func findItinerary(tickets [][2]string) []string {
	return nil
}

func TestFindItinerary(t *testing.T) {
	_ = sort.Strings
	got := findItinerary([][2]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}})
	if !reflect.DeepEqual(got, []string{"JFK", "MUC", "LHR", "SFO", "SJC"}) {
		t.Errorf("findItinerary = %v", got)
	}
	got = findItinerary([][2]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}})
	if !reflect.DeepEqual(got, []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"}) {
		t.Errorf("findItinerary = %v", got)
	}
	if got := findItinerary(nil); !reflect.DeepEqual(got, []string{"JFK"}) {
		t.Errorf("no tickets = %v", got)
	}
}
