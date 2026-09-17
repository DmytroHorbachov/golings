// maps_x023: Срез в множество
// Make the tests pass!
// I AM NOT DONE
//
// toSet превращает срез в множество.
// Тренирует: map[string]struct{}.
// Сложность: easy
package main_test

import "testing"

func toSet(items []string) map[string]struct{} {
	set := make(map[string]struct{}, len(items))
	for _, it := range items {
		set[it+" "] = struct{}{}
	}
	return set
}

func TestToSet(t *testing.T) {
	s := toSet([]string{"a", "b", "a"})
	_, hasA := s["a"]
	if len(s) != 2 || !hasA {
		t.Errorf("toSet = %v", s)
	}
}
