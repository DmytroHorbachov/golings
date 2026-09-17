// slices88
// Make the tests pass!

// I AM NOT DONE
//
// zip объединяет имена и возрасты в срез структур; длина — по короткому срезу.
// Тренирует: параллельный обход двух срезов.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func zip(names []string, ages []int) []Person {
	out := make([]Person, len(names))
	for i := range names {
		out[i] = Person{names[i], ages[i]}
	}
	return out
}

func TestZip(t *testing.T) {
	got := zip([]string{"ann", "bob", "cid"}, []int{30, 40})
	if !reflect.DeepEqual(got, []Person{{"ann", 30}, {"bob", 40}}) {
		t.Errorf("zip = %v", got)
	}
}
