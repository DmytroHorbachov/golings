// slices85
// Make the tests pass!

// I AM NOT DONE
//
// median считает медиану, но после вызова порядок данных вызывающего меняется.
// Тренирует: sort.Ints сортирует срез на месте, в том числе у вызывающего.
// Сложность: hard
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func median(s []int) float64 {
	c := s
	sort.Ints(c)
	n := len(c)
	if n%2 == 1 {
		return float64(c[n/2])
	}
	return float64(c[n/2-1]+c[n/2]) / 2
}

func TestMedian(t *testing.T) {
	data := []int{5, 1, 4, 2}
	if got := median(data); got != 3 {
		t.Errorf("median = %v, want 3", got)
	}
	if !reflect.DeepEqual(data, []int{5, 1, 4, 2}) {
		t.Errorf("data reordered: %v", data)
	}
}
