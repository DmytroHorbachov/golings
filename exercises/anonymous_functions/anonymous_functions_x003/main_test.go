// anonymous_functions_x003: Сортировка по длине
// Make the tests pass!
// I AM NOT DONE
//
// byLength сортирует слова по длине с помощью литерала в sort.Slice.
// Тренирует: функциональный литерал как аргумент.
// Сложность: easy
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func byLength(words []string) {
	sort.Slice(words, func(i, j int) bool {
		return words[i] < words[j]
	})
}

func TestByLength(t *testing.T) {
	w := []string{"banana", "fig", "apple"}
	byLength(w)
	if !reflect.DeepEqual(w, []string{"fig", "apple", "banana"}) {
		t.Errorf("byLength = %v", w)
	}
}
