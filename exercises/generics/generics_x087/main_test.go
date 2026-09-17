// generics_x087: Ограничение как тип переменной
// Make the tests pass!
// I AM NOT DONE
//
// Код пытается объявить срез значений ограничения Number.
// Не компилируется: интерфейс с набором типов можно использовать только как ограничение.
// Тренирует: разница между ограничением и обычным интерфейсом.
// Сложность: hard
package main_test

import "testing"

type Number interface{ ~int | ~float64 }

func Largest(vals []Number) Number {
	m := vals[0]
	for _, v := range vals {
		if v > m {
			m = v
		}
	}
	return m
}

func TestLargest(t *testing.T) {
	if Largest([]int{3, 8, 1}) != 8 || Largest([]float64{0.5, 0.25}) != 0.5 {
		t.Errorf("Largest works incorrectly")
	}
}
