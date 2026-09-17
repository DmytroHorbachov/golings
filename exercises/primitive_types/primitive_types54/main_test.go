// primitive_types54
// Make the tests pass!

// I AM NOT DONE
//
// sortNames сортирует имена по алфавиту без учёта регистра.
// Сейчас все заглавные буквы идут раньше строчных.
// Тренирует: строки сравниваются побайтово: 'Z' (90) < 'a' (97).
// Сложность: hard
package main_test

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

func sortNames(names []string) {
	sort.Slice(names, func(i, j int) bool {
		return names[i] < names[j]
	})
}

func TestSortNames(t *testing.T) {
	_ = strings.ToLower
	names := []string{"bob", "Alice", "carol", "Dave"}
	sortNames(names)
	if want := []string{"Alice", "bob", "carol", "Dave"}; !reflect.DeepEqual(names, want) {
		t.Errorf("sortNames = %v, want %v", names, want)
	}
}
