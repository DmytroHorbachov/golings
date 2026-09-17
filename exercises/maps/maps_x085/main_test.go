// maps_x085: Дробные ключи
// Make the tests pass!
// I AM NOT DONE
//
// Корзины цен ключуются суммой в рублях (float64). Сумма 0.1+0.2
// не находит корзину 0.3.
// Тренирует: вычисленные float-ключи не совпадают с литералами.
// Сложность: hard
package main_test

import (
	"math"
	"testing"
)

type Buckets map[float64]string

func (b Buckets) Add(price float64, name string) { b[price] = name }
func (b Buckets) Get(price float64) string       { return b[price] }

func TestBuckets(t *testing.T) {
	_ = math.Round
	b := Buckets{}
	b.Add(0.3, "cheap")
	b.Add(0.29, "cheaper")
	a, c := 0.1, 0.2
	if got := b.Get(a + c); got != "cheap" {
		t.Errorf("bucket for 0.1+0.2 = %q, want cheap", got)
	}
	if got := b.Get(0.29); got != "cheaper" {
		t.Errorf("bucket for 0.29 = %q", got)
	}
}
