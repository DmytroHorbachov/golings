// functions_x053: Сортировка по двум ключам
// Make the tests pass!
// I AM NOT DONE
//
// sortPeople должна упорядочить людей по возрасту, а при равном возрасте — по имени.
// Тренирует: функцию сравнения с несколькими критериями.
// Сложность: medium
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func sortPeople(ps []Person) {
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].Name < ps[j].Name
	})
}

func TestSortPeople(t *testing.T) {
	ps := []Person{{"Zoe", 30}, {"Adam", 30}, {"Bob", 25}}
	sortPeople(ps)
	want := []Person{{"Bob", 25}, {"Adam", 30}, {"Zoe", 30}}
	if !reflect.DeepEqual(ps, want) {
		t.Errorf("sortPeople = %v, want %v", ps, want)
	}
}
