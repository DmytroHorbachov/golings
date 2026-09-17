// generics_x091: Ноль как «не найдено»
// Make the tests pass!
// I AM NOT DONE
//
// FindMax возвращает максимум, а для пустого среза — нулевое значение.
// Вызывающий код не может отличить пустой срез от среза с максимумом 0.
// Тренирует: нулевое значение T не подходит как признак отсутствия.
// Сложность: hard
package main_test

import "testing"

type Ordered interface{ ~int | ~float64 }

func FindMax[T Ordered](s []T) T {
	var m T
	for i, v := range s {
		if i == 0 || v > m {
			m = v
		}
	}
	return m
}

func summarize(s []int) string {
	if FindMax(s) == 0 {
		return "empty"
	}
	return "has data"
}

func TestSummarize(t *testing.T) {
	if summarize(nil) != "empty" || summarize([]int{-5, 0}) != "has data" {
		t.Errorf("summarize: %q %q", summarize(nil), summarize([]int{-5, 0}))
	}
}
