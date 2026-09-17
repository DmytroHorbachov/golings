// maps_x053: Операции над множествами
// Make the tests pass!
// I AM NOT DONE
//
// union и intersection работают с множествами map[string]struct{}.
// Тренирует: set-операции на map.
// Сложность: medium
package main_test

import "testing"

type Set map[string]struct{}

func union(a, b Set) Set {
	out := Set{}
	for k := range a {
		out[k] = struct{}{}
	}
	return out
}

func intersection(a, b Set) Set {
	out := Set{}
	for k := range a {
		out[k] = struct{}{}
	}
	return out
}

func TestSets(t *testing.T) {
	a := Set{"x": {}, "y": {}}
	b := Set{"y": {}, "z": {}}
	if u := union(a, b); len(u) != 3 {
		t.Errorf("union = %v", u)
	}
	i := intersection(a, b)
	if _, ok := i["y"]; !ok || len(i) != 1 {
		t.Errorf("intersection = %v", i)
	}
}
