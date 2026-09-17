// generics60
// Make the tests pass!

// I AM NOT DONE
//
// SortBy сортирует срез по ключу, вычисленному функцией, не меняя вход.
// Тренирует: обобщённая обёртка над sort.Slice.
// Сложность: medium
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

type Ordered interface{ ~int | ~string }

func SortBy[T any, K Ordered](s []T, key func(T) K) []T {
	sort.Slice(s, func(i, j int) bool { return key(s[i]) > key(s[j]) })
	return s
}

func TestSortBy(t *testing.T) {
	in := []string{"ccc", "a", "bb"}
	got := SortBy(in, func(s string) int { return len(s) })
	if !reflect.DeepEqual(got, []string{"a", "bb", "ccc"}) || in[0] != "ccc" {
		t.Errorf("got %v, input %v", got, in)
	}
}
