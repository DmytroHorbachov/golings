// maps_x055: Корзины по диапазонам
// Make the tests pass!
// I AM NOT DONE
//
// bucketize раскладывает зарплаты по диапазонам: "<50k", "50k-100k", ">=100k".
// Тренирует: вычисление ключа map по значению.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

func label(s int) string {
	switch {
	case s <= 50000:
		return "<50k"
	case s <= 100000:
		return "50k-100k"
	}
	return ">=100k"
}

func bucketize(salaries []int) map[string]int {
	m := map[string]int{}
	for _, s := range salaries {
		m[label(s)]++
	}
	return m
}

func TestBucketize(t *testing.T) {
	got := bucketize([]int{30000, 50000, 99999, 100000, 250000})
	if !reflect.DeepEqual(got, map[string]int{"<50k": 1, "50k-100k": 2, ">=100k": 2}) {
		t.Errorf("bucketize = %v", got)
	}
}
